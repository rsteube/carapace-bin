package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to merge requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_subscribeCmd).Standalone()
	mrCmd.AddCommand(mr_subscribeCmd)

	carapace.Gen(mr_subscribeCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_subscribeCmd, ""),
	)
}
