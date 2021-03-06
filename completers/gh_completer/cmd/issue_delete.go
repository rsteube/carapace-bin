package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_deleteCmd = &cobra.Command{
	Use:   "delete {<number> | <url>}",
	Short: "Delete issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	issueCmd.AddCommand(issue_deleteCmd)

	carapace.Gen(issue_deleteCmd).PositionalCompletion(
		action.ActionIssues(issue_deleteCmd, action.IssueOpts{Open: true, Closed: true}),
	)
}
