package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of relevant issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_statusCmd).Standalone()
	issue_statusCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	issue_statusCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	issue_statusCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	issueCmd.AddCommand(issue_statusCmd)

	carapace.Gen(issue_statusCmd).FlagCompletion(carapace.ActionMap{
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionIssueFields().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
