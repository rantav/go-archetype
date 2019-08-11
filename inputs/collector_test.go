package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectUserInputs(t *testing.T) {
	assert := assert.New(t)

	err := CollectUserInputs(&mockInputsCollector{})
	require.NoError(t, err)

	p := &mockPrompter{
		ID:     "1",
		Answer: "2",
	}
	c := &mockInputsCollector{prompters: []Prompter{p}}
	err = CollectUserInputs(c)
	require.NoError(t, err)
	assert.Equal("2", c.responses["1"].Answer)
}

type mockPrompter struct {
	ID     string
	Answer string
}

func (p *mockPrompter) GetID() string {
	return p.ID
}

func (p *mockPrompter) Prompt() (PromptResponse, error) {
	return p.SetStringResponse(p.Answer), nil
}

func (p *mockPrompter) SetStringResponse(response string) PromptResponse {
	return PromptResponse{
		InputSpec: InputSpec{
			ID: p.ID,
		},
		Answer:   response,
		Answered: true,
	}
}

type mockInputsCollector struct {
	prompters []Prompter
	responses map[string]PromptResponse
}

func (c *mockInputsCollector) GetInputPrompters() []Prompter {
	return c.prompters
}

func (c *mockInputsCollector) SetResponse(r PromptResponse) {
	if c.responses == nil {
		c.responses = make(map[string]PromptResponse)
	}
	c.responses[r.ID] = r
}
