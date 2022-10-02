package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout to an open merge request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_checkoutCmd).Standalone()
	mr_checkoutCmd.Flags().StringP("branch", "b", "", "checkout merge request with <branch> name")
	mr_checkoutCmd.Flags().StringP("set-upstream-to", "u", "", "set tracking of checked out branch to [REMOTE/]BRANCH")
	mr_checkoutCmd.Flags().BoolP("track", "t", false, "set checked out branch to track remote branch, adds remote if needed")
	mrCmd.AddCommand(mr_checkoutCmd)

	carapace.Gen(mr_checkoutCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(mr_checkoutCmd),
		// TODO "set-upstream-to":
	})

	carapace.Gen(mr_checkoutCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_checkoutCmd, ""),
	)
}
