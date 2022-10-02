package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_subscribeCmd).Standalone()
	issueCmd.AddCommand(issue_subscribeCmd)

	carapace.Gen(issue_subscribeCmd).PositionalCompletion(
		action.ActionIssues(issue_subscribeCmd, "opened"),
	)
}
