package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke approval on a merge request <id>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_revokeCmd).Standalone()
	mrCmd.AddCommand(mr_revokeCmd)

	carapace.Gen(mr_revokeCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_revokeCmd, ""),
	)
}
