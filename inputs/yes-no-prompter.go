package inputs

import (
	"fmt"
	"log"
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
	return p.SetStringResponse(answer), nil
}

func (p *yesNoPrompter) GetID() string {
	return p.ID
}

func (p *yesNoPrompter) SetStringResponse(answer string) PromptResponse {
	answer = p.beNiceAndTryToConvert(answer)
	b, err := strconv.ParseBool(answer)
	if err != nil {
		log.Fatalf("Unknown input to yes/no boolean input (use true/false): %s", err)
	}
	if b {
		p.Answer = trueS
	} else {
		p.Answer = "" // This evaluates to false by go tempalates
	}
	p.Answered = true
	return p.PromptResponse
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
