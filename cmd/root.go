package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-archetype",
	Short: "A templating tool",
	Long:  `Use go-archetype to transform project archetypes into existing live projects`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute the current command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
