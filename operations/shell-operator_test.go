package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewShellOperator(t *testing.T) {
	assert := assert.New(t)

	o := newShellOperator(OperationSpec{Sh: []string{"hello"}})
	assert.NotNil(o)
	assert.IsType(&shellOperation{}, o)
}

func TestShellOperatorTemplate(t *testing.T) {
	assert := assert.New(t)

	o := newShellOperator(OperationSpec{Sh: []string{"hello {{.source}}"}})
	require.NotNil(t, o)
	vars := map[string]string{
		"source": "world",
	}
	err := o.Template(vars)
	require.NoError(t, err)

	assert.Equal("hello world", o.sh[0])
}

func TestShellOperatorOperate(t *testing.T) {
	o := newShellOperator(OperationSpec{Sh: []string{"echo hello"}})
	require.NotNil(t, o)
	err := o.Operate()
	require.NoError(t, err)
}

func TestShellOperatorOperateMultiline(t *testing.T) {
	o := newShellOperator(OperationSpec{Sh: []string{"echo hello\necho world"}})
	require.NotNil(t, o)
	err := o.Operate()
	require.NoError(t, err)
}

func TestShellOperatorOperateAndFail(t *testing.T) {
	o := newShellOperator(OperationSpec{Sh: []string{"no-such-command really"}})
	require.NotNil(t, o)
	err := o.Operate()
	require.Error(t, err)
}

func TestShellOperatorOperateMultilineFail(t *testing.T) {
	o := newShellOperator(OperationSpec{Sh: []string{"no-such-command at all\necho well"}})
	require.NotNil(t, o)
	err := o.Operate()
	require.Error(t, err)
}

func TestShellOperatorOperateMultilineFail2(t *testing.T) {
	o := newShellOperator(OperationSpec{Sh: []string{"echo ok\nno-such-command at all"}})
	require.NotNil(t, o)
	err := o.Operate()
	require.Error(t, err)
}
