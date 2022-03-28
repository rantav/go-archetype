package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewShellOperator(t *testing.T) {
	assert := assert.New(t)

	o := newShellOperator(OperationSpec{Sh: []shellOperationSpec{{Cmd: "hello"}}})
	assert.NotNil(o)
	assert.IsType(&shellOperation{}, o)
}

func TestShellOperatorTemplate(t *testing.T) {
	assert := assert.New(t)

	o := newShellOperator(OperationSpec{Sh: []shellOperationSpec{{Cmd: "hello {{.source}}"}}})
	require.NotNil(t, o)
	vars := map[string]string{
		"source": "world",
	}
	err := o.Template(vars)
	require.NoError(t, err)

	assert.Equal("hello world", o.sh[0].cmd)
}

func TestShellOperatorOperate(t *testing.T) {
	t.Run("Single line cmds", func(t *testing.T) {
		t.Run("Successfully execute command", func(t *testing.T) {
			o := newShellOperator(OperationSpec{Sh: []shellOperationSpec{{Cmd: "echo hello"}}})
			require.NotNil(t, o)
			err := o.Operate()
			require.NoError(t, err)
		})

		t.Run("Fail to execute command", func(t *testing.T) {
			o := newShellOperator(OperationSpec{Sh: []shellOperationSpec{{Cmd: "no-such-command really"}}})
			require.NotNil(t, o)
			err := o.Operate()
			require.Error(t, err)
		})
	})

	t.Run("Multiline cmds", func(t *testing.T) {
		t.Run("Multiline shell, split into separate cmds", func(t *testing.T) {
			t.Run("Successfully run cmds", func(t *testing.T) {
				o := newShellOperator(OperationSpec{Sh: []shellOperationSpec{{Cmd: "echo hello\necho world"}}})
				require.NotNil(t, o)
				err := o.Operate()
				require.NoError(t, err)
			})

			t.Run("Fail, the first command doesn't exist", func(t *testing.T) {
				o := newShellOperator(OperationSpec{Sh: []shellOperationSpec{{Cmd: "no-such-command at all\necho well"}}})
				require.NotNil(t, o)
				err := o.Operate()
				require.Error(t, err)
			})

			t.Run("Fail, the second command doesn't exist", func(t *testing.T) {
				o := newShellOperator(OperationSpec{Sh: []shellOperationSpec{{Cmd: "echo ok\nno-such-command at all"}}})
				require.NotNil(t, o)
				err := o.Operate()
				require.Error(t, err)
			})
		})

		t.Run("Multiline shell run as-is", func(t *testing.T) {
			t.Run("Successfully run multiline cmd", func(t *testing.T) {
				o := newShellOperator(
					OperationSpec{
						Sh: []shellOperationSpec{
							{
								Cmd: `
if [ 1 == 1 ]; then
	echo world
fi`,
								Multiline: true,
							},
						},
					},
				)
				require.NotNil(t, o)
				err := o.Operate()
				require.NoError(t, err)
			})
		})
	})
}
