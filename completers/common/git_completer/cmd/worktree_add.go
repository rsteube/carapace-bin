package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var worktree_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a worktree at <path> and checkout <commit-ish> into it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_addCmd).Standalone()

	worktree_addCmd.Flags().StringS("B", "B", "", "create or reset a branch")
	worktree_addCmd.Flags().StringS("b", "b", "", "create a new branch")
	worktree_addCmd.Flags().Bool("checkout", false, "populate the new working tree")
	worktree_addCmd.Flags().BoolP("detach", "d", false, "detach HEAD at named commit")
	worktree_addCmd.Flags().BoolP("force", "f", false, "checkout <branch> even if already checked out in other worktree")
	worktree_addCmd.Flags().Bool("guess-remote", false, "try to match the new branch name with a remote-tracking branch")
	worktree_addCmd.Flags().Bool("lock", false, "keep the new working tree locked")
	worktree_addCmd.Flags().BoolP("quiet", "q", false, "suppress progress reporting")
	worktree_addCmd.Flags().String("reason", "", "reason for locking")
	worktree_addCmd.Flags().Bool("track", false, "set up tracking mode")
	worktreeCmd.AddCommand(worktree_addCmd)

	carapace.Gen(worktree_addCmd).FlagCompletion(carapace.ActionMap{
		"B": git.ActionRefs(git.RefOption{LocalBranches: true}),
		"b": git.ActionRefs(git.RefOption{LocalBranches: true}),
	})

	carapace.Gen(worktree_addCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
