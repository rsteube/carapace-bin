package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/rsteube/carapace-bin/pkg/styles"
	"github.com/spf13/cobra"
)

var pr_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List pull requests in a repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_listCmd).Standalone()
	pr_listCmd.Flags().String("app", "", "Filter by GitHub App author")
	pr_listCmd.Flags().StringP("assignee", "a", "", "Filter by assignee")
	pr_listCmd.Flags().StringP("author", "A", "", "Filter by author")
	pr_listCmd.Flags().StringP("base", "B", "", "Filter by base branch")
	pr_listCmd.Flags().BoolP("draft", "d", false, "Filter by draft state")
	pr_listCmd.Flags().StringP("head", "H", "", "Filter by head branch")
	pr_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	pr_listCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	pr_listCmd.Flags().StringSliceP("label", "l", []string{}, "Filter by label")
	pr_listCmd.Flags().IntP("limit", "L", 30, "Maximum number of items to fetch")
	pr_listCmd.Flags().StringP("search", "S", "", "Search pull requests with `query`")
	pr_listCmd.Flags().StringP("state", "s", "open", "Filter by state: {open|closed|merged|all}")
	pr_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	pr_listCmd.Flags().BoolP("web", "w", false, "List pull requests in the web browser")
	prCmd.AddCommand(pr_listCmd)

	// TODO app completion
	carapace.Gen(pr_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee": action.ActionAssignableUsers(pr_listCmd),
		"author":   gh.ActionUsers(gh.HostOpts{}),
		"base":     action.ActionBranches(pr_listCmd),
		"head":     action.ActionBranches(pr_listCmd),
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPullRequestFields().Invoke(c).Filter(c.Parts).ToA()
		}),
		"label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(pr_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"state": carapace.ActionValues("open", "closed", "merged", "all").StyleF(styles.Gh.ForState),
	})
}
