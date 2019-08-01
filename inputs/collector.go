package inputs

type inputsCollector interface {
	GetInputPrompters() []Prompter
	SetResponse(PromptResponse)
}

func CollectUserInputs(collector inputsCollector) error {
	for _, input := range collector.GetInputPrompters() {
		response, err := input.Prompt()
		if err != nil {
			return err
		}
		collector.SetResponse(response)
	}
	return nil
}
