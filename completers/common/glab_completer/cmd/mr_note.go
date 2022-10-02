package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Add a comment or note to merge request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_noteCmd).Standalone()
	mr_noteCmd.Flags().StringP("message", "m", "", "Comment/Note message")
	mrCmd.AddCommand(mr_noteCmd)

	carapace.Gen(mr_noteCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_noteCmd, ""),
	)
}
