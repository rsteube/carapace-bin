package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var search_reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Search for repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_reposCmd).Standalone()
	search_reposCmd.Flags().Bool("archived", false, "Filter based on archive state")
	search_reposCmd.Flags().String("created", "", "Filter based on created at `date`")
	search_reposCmd.Flags().String("followers", "", "Filter based on `number` of followers")
	search_reposCmd.Flags().String("forks", "", "Filter on `number` of forks")
	search_reposCmd.Flags().String("good-first-issues", "", "Filter on `number` of issues with the 'good first issue' label")
	search_reposCmd.Flags().String("help-wanted-issues", "", "Filter on `number` of issues with the 'help wanted' label")
	search_reposCmd.Flags().String("include-forks", "", "Include forks in fetched repositories: {false|true|only}")
	search_reposCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	search_reposCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	search_reposCmd.Flags().String("language", "", "Filter based on the coding language")
	search_reposCmd.Flags().StringSlice("license", []string{}, "Filter based on license type")
	search_reposCmd.Flags().IntP("limit", "L", 30, "Maximum number of repositories to fetch")
	search_reposCmd.Flags().StringSlice("match", []string{}, "Restrict search to specific field of repository: {name|description|readme}")
	search_reposCmd.Flags().String("number-topics", "", "Filter on `number` of topics")
	search_reposCmd.Flags().String("order", "desc", "Order of repositories returned, ignored unless '--sort' flag is specified: {asc|desc}")
	search_reposCmd.Flags().String("owner", "", "Filter on owner")
	search_reposCmd.Flags().String("size", "", "Filter on a size range, in kilobytes")
	search_reposCmd.Flags().String("sort", "best-match", "Sort fetched repositories: {forks|help-wanted-issues|stars|updated}")
	search_reposCmd.Flags().String("stars", "", "Filter on `number` of stars")
	search_reposCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	search_reposCmd.Flags().StringSlice("topic", []string{}, "Filter on topic")
	search_reposCmd.Flags().String("updated", "", "Filter on last updated at `date`")
	search_reposCmd.Flags().StringSlice("visibility", []string{}, "Filter based on visibility: {public|private|internal}")
	search_reposCmd.Flags().BoolP("web", "w", false, "Open the search query in the web browser")
	searchCmd.AddCommand(search_reposCmd)

	carapace.Gen(search_reposCmd).FlagCompletion(carapace.ActionMap{
		"created":       action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"include-forks": carapace.ActionValues("false", "true", "only"),
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSearchRepositoryFields().Invoke(c).Filter(c.Parts).ToA()
		}),
		"language": action.ActionLanguages(),
		"license": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return gh.ActionLicenses(gh.HostOpts{}).Invoke(c).Filter(c.Parts).ToA()
		}),
		"match": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("name", "description", "readme").Invoke(c).Filter(c.Parts).ToA()
		}),
		"order": carapace.ActionValues("asc", "desc"),
		"owner": gh.ActionOwners(gh.HostOpts{}),
		"sort":  carapace.ActionValues("forks", "help-wanted-issues", "stars", "updated"),
		"topic": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionTopicSearch(search_reposCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"updated": action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"visibility": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("public", "private", "internal").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
