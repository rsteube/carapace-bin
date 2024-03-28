package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_source_gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Reconcile a GitRepository source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_source_gitCmd).Standalone()
	reconcile_sourceCmd.AddCommand(reconcile_source_gitCmd)
}
