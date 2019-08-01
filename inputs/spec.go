package inputs

type InputSpec struct {
	ID   string `yaml:"id"`
	Text string `yaml:"text"`
	Type string `yaml:"type"`
}

func FromSpec(specs []InputSpec) []Prompter {
	var prompters []Prompter
	for _, spec := range specs {
		prompters = append(prompters, NewPrompt(spec))
	}
	return prompters
}
