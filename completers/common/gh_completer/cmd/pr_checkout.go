package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Check out a pull request in git",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_checkoutCmd).Standalone()
	pr_checkoutCmd.Flags().StringP("branch", "b", "", "Local branch name to use (default: the name of the head branch)")
	pr_checkoutCmd.Flags().Bool("detach", false, "Checkout PR with a detached HEAD")
	pr_checkoutCmd.Flags().BoolP("force", "f", false, "Reset the existing local branch to the latest state of the pull request")
	pr_checkoutCmd.Flags().Bool("recurse-submodules", false, "Update all submodules after checkout")
	prCmd.AddCommand(pr_checkoutCmd)

	carapace.Gen(pr_checkoutCmd).PositionalCompletion(
		action.ActionPullRequests(pr_checkoutCmd, action.PullRequestOpts{Open: true}),
	)
}
