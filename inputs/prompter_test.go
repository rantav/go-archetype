package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrompt(t *testing.T) {
	assert := assert.New(t)

	assert.IsType(&yesNoPrompter{}, NewPrompt(InputSpec{Type: "yesno"}))
	assert.IsType(&simpleTextPrompter{}, NewPrompt(InputSpec{Type: "text"}))
}
