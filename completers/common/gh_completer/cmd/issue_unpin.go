package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_unpinCmd = &cobra.Command{
	Use:   "unpin",
	Short: "Unpin a issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_unpinCmd).Standalone()
	issueCmd.AddCommand(issue_unpinCmd)

	carapace.Gen(issue_unpinCmd).PositionalCompletion(
		action.ActionPinnedIssues(issue_unpinCmd),
	)
}
