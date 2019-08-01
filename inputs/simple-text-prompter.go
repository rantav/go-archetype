package inputs

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/pkg/errors"
)

type simpleTextPrompter struct {
	InputSpec
}

func newSimpleTextPrompter(spec InputSpec) Prompter {
	return simpleTextPrompter{InputSpec: spec}
}

func (p simpleTextPrompter) Prompt() (PromptResponse, error) {
	var answer string
	prompt := &survey.Input{
		Message: p.Text,
	}
	err := survey.AskOne(prompt, &answer)
	if err != nil {
		return PromptResponse{}, errors.Wrap(err, "prompt error")
	}
	return PromptResponse{
		InputSpec: p.InputSpec,
		Answer:    answer,
	}, nil
}
