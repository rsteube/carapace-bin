package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_source_gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Suspend reconciliation of a GitRepository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_source_gitCmd).Standalone()
	suspend_sourceCmd.AddCommand(suspend_source_gitCmd)
}
