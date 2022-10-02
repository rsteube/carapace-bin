package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Add a comment to an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_commentCmd).Standalone()
	issue_commentCmd.Flags().StringP("body", "b", "", "The comment body `text`")
	issue_commentCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	issue_commentCmd.Flags().BoolP("editor", "e", false, "Skip prompts and open the text editor to write the body in")
	issue_commentCmd.Flags().BoolP("web", "w", false, "Open the web browser to write the comment")
	issueCmd.AddCommand(issue_commentCmd)

	carapace.Gen(issue_commentCmd).FlagCompletion(carapace.ActionMap{
		"body":      action.ActionKeywordLinks(issue_commentCmd),
		"body-file": carapace.ActionFiles(),
	})
	carapace.Gen(issue_commentCmd).PositionalCompletion(
		action.ActionIssues(issue_commentCmd, action.IssueOpts{Open: true}),
	)
}
