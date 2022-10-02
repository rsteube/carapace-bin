package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Close an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_closeCmd).Standalone()
	issueCmd.AddCommand(issue_closeCmd)

	carapace.Gen(issue_closeCmd).PositionalCompletion(
		action.ActionIssues(issue_closeCmd, "opened"),
	)
}
