package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer issue to another repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_transferCmd).Standalone()
	issueCmd.AddCommand(issue_transferCmd)

	carapace.Gen(issue_transferCmd).PositionalCompletion(
		action.ActionIssues(issue_transferCmd, action.IssueOpts{Open: true, Closed: true}),
		action.ActionOwnerRepositories(issue_transferCmd),
	)
}
