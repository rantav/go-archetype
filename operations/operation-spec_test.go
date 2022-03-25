package operations

import (
	"testing"

	"github.com/rantav/go-archetype/log"
	"github.com/stretchr/testify/assert"
)

func TestFromSpec(t *testing.T) {
	assert := assert.New(t)

	o := FromSpec(Spec{
		Operations: []OperationSpec{
			{
				Sh: []string{"hello"},
			},
		}}, log.NopLogger{})
	assert.NotNil(o)
	assert.Len(o, 1)
	assert.IsType(&shellOperation{}, o[0])
}
