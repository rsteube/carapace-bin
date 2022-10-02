package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var worktree_unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock a worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_unlockCmd).Standalone()

	worktreeCmd.AddCommand(worktree_unlockCmd)
}
