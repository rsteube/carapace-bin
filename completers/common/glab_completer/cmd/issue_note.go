package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Add a comment or note to an issue on GitLab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_noteCmd).Standalone()
	issue_noteCmd.Flags().StringP("message", "m", "", "Comment/Note message")
	issueCmd.AddCommand(issue_noteCmd)

	carapace.Gen(issue_noteCmd).PositionalCompletion(
		action.ActionIssues(issue_noteCmd, "opened"),
	)
}
