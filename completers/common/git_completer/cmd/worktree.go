package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var worktreeCmd = &cobra.Command{
	Use:   "worktree",
	Short: "Manage multiple working trees",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktreeCmd).Standalone()

	rootCmd.AddCommand(worktreeCmd)
}
