package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete stale references",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_pruneCmd).Standalone()

	remote_pruneCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	remoteCmd.AddCommand(remote_pruneCmd)

	carapace.Gen(remote_pruneCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
