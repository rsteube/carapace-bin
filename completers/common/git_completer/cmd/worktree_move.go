package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var worktree_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a worktree to a new location",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_moveCmd).Standalone()

	worktree_moveCmd.Flags().BoolP("force", "f", false, "force move even if worktree is dirty or locked")
	worktreeCmd.AddCommand(worktree_moveCmd)

	carapace.Gen(worktree_moveCmd).PositionalCompletion(
		git.ActionWorktrees(),
		carapace.ActionDirectories(),
	)
}
