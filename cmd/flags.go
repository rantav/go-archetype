package cmd

import "github.com/spf13/cobra"

func addRequiredStringFlag(command *cobra.Command, name, value, usage string) *string {
	ref := command.Flags().String(name, value, usage)
	err := command.MarkFlagRequired(name)
	if err != nil {
		panic(err)
	}
	return ref
}
