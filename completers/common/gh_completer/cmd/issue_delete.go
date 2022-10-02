package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_deleteCmd).Standalone()
	issue_deleteCmd.Flags().Bool("confirm", false, "confirm deletion without prompting")
	issueCmd.AddCommand(issue_deleteCmd)

	carapace.Gen(issue_deleteCmd).PositionalCompletion(
		action.ActionIssues(issue_deleteCmd, action.IssueOpts{Open: true, Closed: true}),
	)
}
