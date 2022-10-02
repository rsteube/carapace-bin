package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var worktree_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Prune worktree information in $GIT_DIR/worktrees",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_pruneCmd).Standalone()

	worktree_pruneCmd.Flags().BoolP("dry-run", "n", false, "do not remove, show only")
	worktree_pruneCmd.Flags().String("expire", "", "expire working trees older than <time>")
	worktree_pruneCmd.Flags().BoolP("verbose", "v", false, "report pruned working trees")
	worktreeCmd.AddCommand(worktree_pruneCmd)
}
