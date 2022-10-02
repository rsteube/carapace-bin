package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge/Accept merge requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_mergeCmd).Standalone()
	mr_mergeCmd.Flags().StringP("message", "m", "", "Custom merge commit message")
	mr_mergeCmd.Flags().BoolP("rebase", "r", false, "Rebase the commits onto the base branch")
	mr_mergeCmd.Flags().BoolP("remove-source-branch", "d", false, "Remove source branch on merge")
	mr_mergeCmd.Flags().String("sha", "", "Merge Commit sha")
	mr_mergeCmd.Flags().BoolP("squash", "s", false, "Squash commits on merge")
	mr_mergeCmd.Flags().String("squash-message", "", "Custom Squash commit message")
	mr_mergeCmd.Flags().Bool("when-pipeline-succeeds", true, "Merge only when pipeline succeeds")
	mr_mergeCmd.Flags().BoolP("yes", "y", false, "Skip submission confirmation prompt")
	mrCmd.AddCommand(mr_mergeCmd)

	carapace.Gen(mr_mergeCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_mergeCmd, "opened"),
	)
}
