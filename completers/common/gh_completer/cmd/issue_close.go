package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_closeCmd).Standalone()
	issue_closeCmd.Flags().StringP("comment", "c", "", "Leave a closing comment")
	issue_closeCmd.Flags().StringP("reason", "r", "", "Reason for closing: {completed|not planned}")
	issueCmd.AddCommand(issue_closeCmd)

	carapace.Gen(issue_closeCmd).FlagCompletion(carapace.ActionMap{
		"comment": action.ActionKeywordLinks(issue_closeCmd),
		"reason":  carapace.ActionValues("completed", "not planned"),
	})

	carapace.Gen(issue_closeCmd).PositionalCompletion(
		action.ActionIssues(issue_closeCmd, action.IssueOpts{Open: true}),
	)
}
