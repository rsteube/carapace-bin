package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_unsubscribeCmd = &cobra.Command{
	Use:   "unsubscribe",
	Short: "Unsubscribe from merge requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_unsubscribeCmd).Standalone()
	mrCmd.AddCommand(mr_unsubscribeCmd)

	carapace.Gen(mr_unsubscribeCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_unsubscribeCmd, ""),
	)
}
