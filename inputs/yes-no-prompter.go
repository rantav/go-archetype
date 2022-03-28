package inputs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

const (
	trueS  = "true"
	falseS = "false"
)

type yesNoPrompter struct {
	PromptResponse
}

func newYesNoPrompter(spec InputSpec) *yesNoPrompter {
	return &yesNoPrompter{PromptResponse: PromptResponse{InputSpec: spec}}
}

func (p *yesNoPrompter) Prompt() (PromptResponse, error) {
	if p.Answered {
		return p.PromptResponse, nil
	}
	var (
		yes    bool
		answer string
	)
	prompt := &survey.Confirm{
		Message: p.Text,
	}
	err := survey.AskOne(prompt, &yes)
	if err != nil {
		return PromptResponse{}, fmt.Errorf("prompt error: %w", err)
	}
	if yes {
		answer = trueS
	} else {
		answer = falseS
	}
	return p.SetStringResponse(answer)
}

func (p *yesNoPrompter) GetID() string {
	return p.ID
}

func (p *yesNoPrompter) SetStringResponse(answer string) (PromptResponse, error) {
	answer = p.beNiceAndTryToConvert(answer)
	b, err := strconv.ParseBool(answer)
	if err != nil {
		return PromptResponse{}, fmt.Errorf("unknown input to yes/no boolean input (use true/false): %w", err)
	}
	if b {
		p.Answer = trueS
	} else {
		p.Answer = "" // This evaluates to false by go tempalates
	}
	p.Answered = true
	return p.PromptResponse, nil
}

// Tries to find a suitable conversion b/w the input string and a true/false value.
// If not found, just returns the original value itself
func (p *yesNoPrompter) beNiceAndTryToConvert(str string) string {
	switch strings.ToLower(str) {
	case "yes", "ok", "sure", "why not":
		return trueS
	case "", "no", "hell no", "as if":
		return falseS
	default:
		return str
	}
}

var _ Prompter = &yesNoPrompter{}
