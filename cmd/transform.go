package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/rantav/go-archetype/generator"
	"github.com/rantav/go-archetype/log"
)

// CLI flags
var (
	transformationsFile *string
	source              *string
	destination         *string
	logLevel            *string
)

// transformCmd represents the transform command
var transformCmd = &cobra.Command{
	Use:   "transform",
	Short: "Transform a blueprint to a live project",
	Run: func(cmd *cobra.Command, args []string) {
		envLogLevel, ok := os.LookupEnv("LOG_LEVEL")
		if ok {
			logLevel = &envLogLevel
		}

		logger := log.NewZeroLogger(*logLevel)

		err := generator.Generate(*transformationsFile, *source, *destination, args, logger)
		if err != nil {
			logger.Fatalf("error generating: %s", err)
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

	logLevel = transformCmd.Flags().String("log-level", "info", "The minimal level for logging")
}
