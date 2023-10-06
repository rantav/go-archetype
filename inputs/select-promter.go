package inputs

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type selectPrompter struct {
	PromptResponse
}

func newSelectPrompter(spec InputSpec) *selectPrompter {
	return &selectPrompter{PromptResponse: PromptResponse{InputSpec: spec}}
}

func (p *selectPrompter) GetID() string {
	return p.ID
}

func (p *selectPrompter) Prompt() (PromptResponse, error) {
	if p.Answered {
		return p.PromptResponse, nil
	}

	var answer string
	prompt := &survey.Select{
		Message: p.Text,
		Options: p.Options,
		Default: p.Options[0],
	}
	err := survey.AskOne(prompt, &answer)
	if err != nil {
		return PromptResponse{}, fmt.Errorf("prompt error: %w", err)
	}
	return p.SetStringResponse(answer)
}

func (p *selectPrompter) SetStringResponse(answer string) (PromptResponse, error) {
	if err := p.validateAnswer(answer); err != nil {
		return PromptResponse{}, err
	}
	p.Answer = answer
	p.Answered = true
	return p.PromptResponse, nil
}

func (p *selectPrompter) validateAnswer(answer string) error {
	for _, b := range p.Options {
		if b == answer {
			return nil
		}
	}
	return fmt.Errorf("answer %s is not in a list %s", answer, p.Options)
}

var _ Prompter = &selectPrompter{}
