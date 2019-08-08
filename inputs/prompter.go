package inputs

type PromptResponse struct {
	InputSpec
	// Answer string
	Answer string
	// Has this been answered already?
	Answered bool
}

// Prompter is a user-input prompter capabale of storing user inputs
// as well as prompting for interactive user input
type Prompter interface {
	GetID() string
	Prompt() (PromptResponse, error)
	SetStringResponse(string) PromptResponse
}

// Create a new prompt based on spec.Type
func NewPrompt(spec InputSpec) Prompter {
	switch spec.Type {
	case "text":
		return newSimpleTextPrompter(spec)
	case "yesno":
		return newYesNoPrompter(spec)
	default:
		panic("Unknown user input type")
	}
}
