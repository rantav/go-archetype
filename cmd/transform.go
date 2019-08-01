package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"

	"github.com/rantav/go-archetype/inputs"
	"github.com/rantav/go-archetype/transformer"
	"github.com/rantav/go-archetype/types"
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
		transformations, err := transformer.Read(*transformationsFile)
		if err != nil {
			panic(err)
		}
		err = inputs.CollectUserInputs(transformations)
		if err != nil {
			panic(err)
		}

		err = transformations.Template()
		if err != nil {
			panic(err)
		}

		spew.Dump(transformations)

		sourcePath := types.Path(*source)
		destinationPath := types.Path(*destination)
		err = transformer.Transform(sourcePath, destinationPath, *transformations)
		if err != nil {
			panic(err)
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
