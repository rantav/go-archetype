package inputs

import (
	"fmt"

	"github.com/spf13/pflag"
)

// Parse CLI args that were passed to the templating system.
// Argse are parsed and those that have values are stored in the inputsCollector.
func ParseCLIArgsInputs(collector inputsCollector, args []string) error {
	prompters := collector.GetInputPrompters()
	inputNames := collectInputNames(prompters)
	providedCLIArgs, err := parseUserInputCliArgs(inputNames, args)
	if err != nil {
		return err
	}
	for i, input := range prompters {
		provided := providedCLIArgs[i]
		if provided != nil {
			response := input.SetStringResponse(*provided)
			collector.SetResponse(response)
		}
	}
	return nil
}

func collectInputNames(prompters []Prompter) []string {
	var ids []string
	for _, prompter := range prompters {
		ids = append(ids, prompter.GetID())
	}
	return ids
}

// Parses user's input and returns an array the same size as definedArgNames where in each position
// the user's provided value is assigned.
// Where user did not provide the value, the returned *string is nil
func parseUserInputCliArgs(definedArgNames []string, providedArgs []string) ([]*string, error) {
	const notSet = "addje8382372742@#!@#$!@*#$!JFQW(WFQHR#(RNFW(*FWNFIW(™‹adsfqwf" // Just a random string
	flagSet := pflag.NewFlagSet("user inputs", pflag.ContinueOnError)
	providedCLIArgs := make([]*string, len(definedArgNames))
	for i, name := range definedArgNames {
		providedCLIArgs[i] = flagSet.String(name, notSet, "")
	}
	err := flagSet.Parse(providedArgs)
	if err != nil {
		return nil, fmt.Errorf("parsing CLI args: %w", err)
	}
	// clean up the notSet values
	for i := range providedCLIArgs {
		if *providedCLIArgs[i] == notSet {
			providedCLIArgs[i] = nil
		}
	}
	return providedCLIArgs, nil
}
