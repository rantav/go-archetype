package operations

import (
	"testing"

	"github.com/rantav/go-archetype/log"
	"github.com/stretchr/testify/assert"
)

func TestNewOperator(t *testing.T) {
	assert := assert.New(t)

	o := NewOperator(OperationSpec{Sh: []string{"hello"}}, log.NopLogger{})
	assert.NotNil(o)
	assert.IsType(&shellOperation{}, o)
}
