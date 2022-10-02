package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_closeCmd).Standalone()
	pr_closeCmd.Flags().StringP("comment", "c", "", "Leave a closing comment")
	pr_closeCmd.Flags().BoolP("delete-branch", "d", false, "Delete the local and remote branch after close")
	prCmd.AddCommand(pr_closeCmd)

	carapace.Gen(pr_closeCmd).FlagCompletion(carapace.ActionMap{
		"comment": action.ActionKeywordLinks(pr_closeCmd),
	})

	carapace.Gen(pr_closeCmd).PositionalCompletion(
		action.ActionPullRequests(pr_closeCmd, action.PullRequestOpts{Open: true}),
	)
}
