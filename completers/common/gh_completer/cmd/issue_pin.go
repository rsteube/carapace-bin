package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "Pin a issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_pinCmd).Standalone()
	issueCmd.AddCommand(issue_pinCmd)

	carapace.Gen(issue_pinCmd).PositionalCompletion(
		action.ActionIssues(issue_pinCmd, action.IssueOpts{Open: true}),
	)
}
