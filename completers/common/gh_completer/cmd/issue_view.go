package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_viewCmd).Standalone()
	issue_viewCmd.Flags().BoolP("comments", "c", false, "View issue comments")
	issue_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	issue_viewCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	issue_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	issue_viewCmd.Flags().BoolP("web", "w", false, "Open an issue in the browser")
	issueCmd.AddCommand(issue_viewCmd)

	carapace.Gen(issue_viewCmd).FlagCompletion(carapace.ActionMap{
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionIssueFields().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(issue_viewCmd).PositionalCompletion(
		action.ActionIssues(issue_viewCmd, action.IssueOpts{Open: true}),
	)
}
