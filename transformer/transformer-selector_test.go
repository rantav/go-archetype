package transformer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransformer(t *testing.T) {
	assert := assert.New(t)
	assert.IsType(&textReplacer{}, newTransformer(transformationSpec{Type: "replace"}))
	assert.IsType(&includeTransformer{}, newTransformer(transformationSpec{Type: "include"}))
}
