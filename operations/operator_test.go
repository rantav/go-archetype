package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperator(t *testing.T) {
	assert := assert.New(t)

	o := NewOperator(OperationSpec{Sh: []string{"hello"}})
	assert.NotNil(o)
	assert.IsType(&shellOperation{}, o)
}
