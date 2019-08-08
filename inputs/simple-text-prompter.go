package inputs

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/pkg/errors"
)

type simpleTextPrompter struct {
	PromptResponse
}

func newSimpleTextPrompter(spec InputSpec) *simpleTextPrompter {
	return &simpleTextPrompter{PromptResponse: PromptResponse{InputSpec: spec}}
}

func (p *simpleTextPrompter) Prompt() (PromptResponse, error) {
	if p.Answered {
		return p.PromptResponse, nil
	}
	var answer string
	prompt := &survey.Input{
		Message: p.Text,
	}
	err := survey.AskOne(prompt, &answer)
	if err != nil {
		return PromptResponse{}, errors.Wrap(err, "prompt error")
	}
	return p.SetStringResponse(answer), nil
}

func (p *simpleTextPrompter) GetID() string {
	return p.ID
}

func (p *simpleTextPrompter) SetStringResponse(answer string) PromptResponse {
	p.Answer = answer
	p.Answered = true
	return p.PromptResponse
}

var _ Prompter = &simpleTextPrompter{}
