package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests for a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
