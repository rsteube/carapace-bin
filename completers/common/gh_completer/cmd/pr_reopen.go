package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_reopenCmd = &cobra.Command{
	Use:   "reopen",
	Short: "Reopen a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_reopenCmd).Standalone()
	pr_reopenCmd.Flags().StringP("comment", "c", "", "Add a reopening comment")
	prCmd.AddCommand(pr_reopenCmd)

	carapace.Gen(pr_reopenCmd).FlagCompletion(carapace.ActionMap{
		"comment": action.ActionKeywordLinks(pr_reopenCmd),
	})

	carapace.Gen(pr_reopenCmd).PositionalCompletion(
		action.ActionPullRequests(pr_reopenCmd, action.PullRequestOpts{Closed: true}),
	)
}
