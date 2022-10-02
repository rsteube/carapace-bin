package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Add a review to a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_reviewCmd).Standalone()
	pr_reviewCmd.Flags().BoolP("approve", "a", false, "Approve pull request")
	pr_reviewCmd.Flags().StringP("body", "b", "", "Specify the body of a review")
	pr_reviewCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	pr_reviewCmd.Flags().BoolP("comment", "c", false, "Comment on a pull request")
	pr_reviewCmd.Flags().BoolP("request-changes", "r", false, "Request changes on a pull request")
	prCmd.AddCommand(pr_reviewCmd)

	carapace.Gen(pr_reviewCmd).FlagCompletion(carapace.ActionMap{
		"body":      action.ActionKeywordLinks(pr_reviewCmd),
		"body-file": carapace.ActionFiles(),
		"comment":   action.ActionKeywordLinks(pr_reviewCmd),
	})

	carapace.Gen(pr_reviewCmd).PositionalCompletion(
		action.ActionPullRequests(pr_reviewCmd, action.PullRequestOpts{Open: true}),
	)
}
