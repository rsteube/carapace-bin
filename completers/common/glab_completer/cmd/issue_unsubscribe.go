package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_unsubscribeCmd = &cobra.Command{
	Use:   "unsubscribe",
	Short: "Unsubscribe to an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_unsubscribeCmd).Standalone()
	issueCmd.AddCommand(issue_unsubscribeCmd)

	carapace.Gen(issue_unsubscribeCmd).PositionalCompletion(
		action.ActionIssues(issue_unsubscribeCmd, "opened"),
	)
}
