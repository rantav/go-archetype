package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/rantav/go-archetype/generator"
)

// CLI flags
var (
	transformationsFile *string
	source              *string
	destination         *string
)

// transformCmd represents the transform command
var transformCmd = &cobra.Command{
	Use:   "transform",
	Short: "Transform a blueprint to a live project",
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.Generate(*transformationsFile, *source, *destination, args)
		if err != nil {
			log.Fatalf("error generating: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(transformCmd)
	transformationsFile = addRequiredStringFlag(transformCmd, "transformations", "",
		"Location of your transformations.yaml file")
	source = addRequiredStringFlag(transformCmd, "source", ".",
		"Location of the source (blueprint) files")
	destination = addRequiredStringFlag(transformCmd, "destination", "",
		"Location of the destination (generated) files")
}
