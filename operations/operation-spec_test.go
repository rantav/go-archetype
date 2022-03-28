package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/rantav/go-archetype/log"
)

func TestFromSpec(t *testing.T) {
	t.Run("Basic spec as a single line", func(t *testing.T) {
		spec := Spec{
			Operations: []OperationSpec{
				{
					Sh: []shellOperationSpec{{Cmd: "hello"}},
				},
			},
		}

		expected := []Operator{
			&shellOperation{
				sh: []shellCmdOperation{
					{
						cmd:       "hello",
						multiline: false,
					},
				},
				logger: log.NopLogger{},
			},
		}

		actual := FromSpec(spec, log.NopLogger{})

		assert.Equal(t, expected, actual)
	})
}

func TestShellOperationSpec_UnmarshalYAML(t *testing.T) {
	t.Run("Successfully unmarshal YAML", func(t *testing.T) {
		// language=YAML
		yamlFile := `---
operations:
  - sh:
    - echo 1
    - echo 2
    - cmd: |-
        if [ 1 == 1 ]; then
          echo "OK!"
        fi
      multiline: true
    - echo 3
`
		var spec Spec

		err := yaml.Unmarshal([]byte(yamlFile), &spec)
		require.NoError(t, err)

		expected := Spec{
			Operations: []OperationSpec{
				{
					Sh: []shellOperationSpec{
						{
							Cmd:       "echo 1",
							Multiline: false,
						},
						{
							Cmd:       "echo 2",
							Multiline: false,
						},
						{
							Cmd: `if [ 1 == 1 ]; then
  echo "OK!"
fi`,
							Multiline: true,
						},
						{
							Cmd:       "echo 3",
							Multiline: false,
						},
					},
				},
			},
		}

		assert.Equal(t, expected, spec)
	})

	t.Run("Error on unmarshalling", func(t *testing.T) {
		// language=YAML
		yamlFile := `---
operations:
  - sh:
    -
      - echo 1
      - echo 2
`
		var spec Spec

		err := yaml.Unmarshal([]byte(yamlFile), &spec)
		require.Error(t, err)
	})
}
