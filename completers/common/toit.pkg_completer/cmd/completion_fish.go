package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completion_fishCmd = &cobra.Command{
	Use:   "fish",
	Short: "Generate the autocompletion script for fish",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completion_fishCmd).Standalone()
	completion_fishCmd.Flags().Bool("no-descriptions", false, "disable completion descriptions")
	completionCmd.AddCommand(completion_fishCmd)
}
