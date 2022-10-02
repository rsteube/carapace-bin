package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ci_retryCmd = &cobra.Command{
	Use:   "retry",
	Short: "Retry a CI job",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_retryCmd).Standalone()
	ciCmd.AddCommand(ci_retryCmd)

	// TODO positional completion
}
