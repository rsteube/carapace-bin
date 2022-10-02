package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_readyCmd = &cobra.Command{
	Use:   "ready",
	Short: "Mark a pull request as ready for review",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_readyCmd).Standalone()
	pr_readyCmd.Flags().Bool("undo", false, "Convert a pull request to \"draft\"")
	prCmd.AddCommand(pr_readyCmd)

	carapace.Gen(pr_readyCmd).PositionalCompletion(
		action.ActionPullRequests(pr_readyCmd, action.PullRequestOpts{Open: true, DraftOnly: true}),
	)
}
