package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Get issues related to a particular merge request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_issuesCmd).Standalone()
	mrCmd.AddCommand(mr_issuesCmd)

	carapace.Gen(mr_issuesCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_issuesCmd, ""),
	)
}
