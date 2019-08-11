package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSpec(t *testing.T) {
	assert := assert.New(t)

	prompters := FromSpec([]InputSpec{
		{
			ID:   "1",
			Type: "text",
			Text: "what's your name?",
		},
		{
			ID:   "2",
			Type: "yesno",
			Text: "are you sure?",
		},
	})

	assert.Len(prompters, 2)
	assert.IsType(&simpleTextPrompter{}, prompters[0])
	assert.IsType(&yesNoPrompter{}, prompters[1])
}
