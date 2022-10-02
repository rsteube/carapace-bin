package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Automatically rebase the source_branch of the merge request against its target_branch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_rebaseCmd).Standalone()
	mrCmd.AddCommand(mr_rebaseCmd)

	carapace.Gen(mr_rebaseCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_rebaseCmd, ""),
	)
}
