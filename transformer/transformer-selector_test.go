package transformer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rantav/go-archetype/log"
)

func TestNewTransformer(t *testing.T) {
	assert := assert.New(t)
	assert.IsType(&textReplacer{}, newTransformer(transformationSpec{Type: "replace"}, log.NopLogger{}))
	assert.IsType(&includeTransformer{}, newTransformer(transformationSpec{Type: "include"}, log.NopLogger{}))
	assert.IsType(&fileRenamer{}, newTransformer(transformationSpec{Type: "rename"}, log.NopLogger{}))
}
