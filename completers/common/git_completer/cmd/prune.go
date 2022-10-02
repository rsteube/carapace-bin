package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Prune all unreachable objects from the object database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	pruneCmd.Flags().BoolP("dry-run", "n", false, "do not remove, show only")
	pruneCmd.Flags().Bool("exclude-promisor-objects", false, "limit traversal to objects outside promisor packfiles")
	pruneCmd.Flags().String("expire", "", "expire objects older than <time>")
	pruneCmd.Flags().Bool("progress", false, "show progress")
	pruneCmd.Flags().BoolP("verbose", "v", false, "report pruned objects")
	rootCmd.AddCommand(pruneCmd)

	carapace.Gen(pruneCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{LocalBranches: true, Tags: true}),
	)
}
