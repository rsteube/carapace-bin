package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of relevant pull requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_statusCmd).Standalone()
	pr_statusCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	pr_statusCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	pr_statusCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	prCmd.AddCommand(pr_statusCmd)

	carapace.Gen(pr_statusCmd).FlagCompletion(carapace.ActionMap{
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPullRequestFields().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
