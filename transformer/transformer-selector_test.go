package transformer

import (
	"testing"

	"github.com/rantav/go-archetype/log"
	"github.com/stretchr/testify/assert"
)

func TestNewTransformer(t *testing.T) {
	assert := assert.New(t)
	assert.IsType(&textReplacer{}, newTransformer(transformationSpec{Type: "replace"}, log.NopLogger{}))
	assert.IsType(&includeTransformer{}, newTransformer(transformationSpec{Type: "include"}, log.NopLogger{}))
}
