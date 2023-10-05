package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSpec(t *testing.T) {
	assert := assert.New(t)

	prompters := FromSpec([]InputSpec{
		{
			ID:   "0",
			Type: "text",
			Text: "what's your name?",
		},
		{
			ID:   "1",
			Text: "are you sure?",
			Type: "yesno",
		},
		{
			ID:      "2",
			Type:    "select",
			Text:    "Choose value:",
			Options: []string{"one", "two", "three"},
		},
	})

	assert.Len(prompters, 3)
	assert.IsType(&simpleTextPrompter{}, prompters[0])
	assert.IsType(&yesNoPrompter{}, prompters[1])
	assert.IsType(&selectPrompter{}, prompters[2])
}
