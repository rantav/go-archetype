package inputs

import (
	"github.com/spf13/pflag"
)

// Parse CLI args that were passed to the templating system.
// Argse are parsed and those that have values are stored in the inputsCollector.
func ParseCLIArgsInputs(collector inputsCollector, args []string) error {
	flagSet := pflag.NewFlagSet("user inputs", pflag.ContinueOnError)
	prompters := collector.GetInputPrompters()
	providedCLIArgs := make([]*string, len(prompters))
	for i, input := range prompters {
		providedCLIArgs[i] = flagSet.String(input.GetID(), "", "")
	}
	err := flagSet.Parse(args)
	if err != nil {
		return err
	}
	for i, input := range prompters {
		provided := providedCLIArgs[i]
		if provided != nil && *provided != "" {
			response := input.SetStringResponse(*provided)
			collector.SetResponse(response)
		}
	}

	return nil
}
