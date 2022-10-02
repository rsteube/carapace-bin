package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_reopenCmd = &cobra.Command{
	Use:   "reopen",
	Short: "Reopen a closed issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_reopenCmd).Standalone()
	issueCmd.AddCommand(issue_reopenCmd)

	carapace.Gen(issue_reopenCmd).PositionalCompletion(
		action.ActionIssues(issue_reopenCmd, "closed"),
	)
}
