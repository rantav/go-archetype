package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rantav/go-archetype/log"
)

func TestNewOperator(t *testing.T) {
	assert := assert.New(t)

	o := NewOperator(OperationSpec{Sh: []shellOperationSpec{{Cmd: "hello"}}}, log.NopLogger{})
	assert.NotNil(o)
	assert.IsType(&shellOperation{}, o)
}
