package inputs

type PromptResponse struct {
	InputSpec
	Answer string
}

type Prompter interface {
	Prompt() (PromptResponse, error)
}

func NewPrompt(spec InputSpec) Prompter {
	switch spec.Type {
	case "text":
		return newSimpleTextPrompter(spec)
	default:
		panic("Unknown user input type")
	}
}
