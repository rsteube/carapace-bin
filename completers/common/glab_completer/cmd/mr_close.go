package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close merge requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_closeCmd).Standalone()
	mrCmd.AddCommand(mr_closeCmd)

	carapace.Gen(mr_closeCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_closeCmd, ""),
	)
}
