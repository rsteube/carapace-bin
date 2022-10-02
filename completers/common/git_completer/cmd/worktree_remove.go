package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var worktree_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_removeCmd).Standalone()

	worktree_removeCmd.Flags().BoolP("force", "f", false, "force removal even if worktree is dirty or locked")
	worktreeCmd.AddCommand(worktree_removeCmd)

	carapace.Gen(worktree_removeCmd).PositionalCompletion(
		git.ActionWorktrees(),
	)
}
