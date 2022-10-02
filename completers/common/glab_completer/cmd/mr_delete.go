package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete merge requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_deleteCmd).Standalone()
	mrCmd.AddCommand(mr_deleteCmd)

	carapace.Gen(mr_deleteCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_deleteCmd, ""),
	)
}
