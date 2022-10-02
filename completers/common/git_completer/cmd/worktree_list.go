package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var worktree_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List details of each worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_listCmd).Standalone()

	worktree_listCmd.Flags().String("expire", "", "add 'prunable' annotation to worktrees older than <time>")
	worktree_listCmd.Flags().Bool("porcelain", false, "machine-readable output")
	worktree_listCmd.Flags().BoolP("verbose", "v", false, "show extended annotations and reasons, if available")
	worktree_listCmd.Flags().BoolS("z", "z", false, "terminate records with a NUL character")
	worktreeCmd.AddCommand(worktree_listCmd)
}
