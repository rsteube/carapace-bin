package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var worktree_repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "Repair worktree administrative files, if possible",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_repairCmd).Standalone()

	worktreeCmd.AddCommand(worktree_repairCmd)

	carapace.Gen(worktree_repairCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
