package transformer

import (
	"testing"

	"github.com/rantav/go-archetype/types"
	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	assert := assert.New(t)
	tester := func(name string, truthy bool, marker string, input, expectedOutput types.FileContents) {
		transformer := includeTransformer{
			truthy:       truthy,
			regionMarker: marker,
		}
		output := transformer.Transform(input)
		assert.Equalf(string(expectedOutput), string(output), "Test failed: %s", name)
	}
	tests := []struct {
		name   string
		truthy bool
		marker string
		input  string
		output string
	}{
		{
			"truthy, all empty",
			true,
			"",
			"",
			"",
		},
		{
			"falsy, all empty",
			false,
			"",
			"",
			"",
		},
		{
			"truthy, with a marker",
			true,
			"__1__",
			`1
BEGIN __1__
2
END __1__
3
`,
			`1
2
3
`,
		},
		{
			"falsy, with a marker",
			false,
			"__1__",
			`1
BEGIN __1__
2
END __1__
3
`,
			`1
3
`,
		},
	}
	for _, test := range tests {
		tester(test.name, test.truthy, test.marker,
			types.FileContents(test.input), types.FileContents(test.output))
	}
}
