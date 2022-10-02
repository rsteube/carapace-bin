package cmd

import (
	"github.com/spf13/cobra"
)

var completion_fishCmd = &cobra.Command{
	Use:   "fish",
	Short: "generate autocompletion script for fish",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	completion_fishCmd.Flags().Bool("no-descriptions", false, "disable completion descriptions")
	completionCmd.AddCommand(completion_fishCmd)
}
