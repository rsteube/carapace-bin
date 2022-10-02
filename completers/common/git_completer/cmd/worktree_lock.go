package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var worktree_lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock administrative files from being pruned automatically",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_lockCmd).Standalone()

	worktree_lockCmd.Flags().String("expire", "", "add 'prunable' annotation to worktrees older than <time>")
	worktree_lockCmd.Flags().Bool("porcelain", false, "machine-readable output")
	worktree_lockCmd.Flags().BoolP("verbose", "v", false, "show extended annotations and reasons, if available")
	worktree_lockCmd.Flags().BoolS("z", "z", false, "terminate records with a NUL character")
	worktreeCmd.AddCommand(worktree_lockCmd)

	carapace.Gen(worktree_lockCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
