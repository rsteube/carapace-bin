package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_viewCmd).Standalone()
	pr_viewCmd.Flags().BoolP("comments", "c", false, "View pull request comments")
	pr_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	pr_viewCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	pr_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	pr_viewCmd.Flags().BoolP("web", "w", false, "Open a pull request in the browser")
	prCmd.AddCommand(pr_viewCmd)

	carapace.Gen(pr_viewCmd).FlagCompletion(carapace.ActionMap{
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPullRequestFields().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(pr_viewCmd).PositionalCompletion(
		action.ActionPullRequests(pr_viewCmd, action.PullRequestOpts{Open: true}),
	)
}
