package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/rsteube/carapace-bin/pkg/styles"
	"github.com/spf13/cobra"
)

var search_issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Search for issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_issuesCmd).Standalone()
	search_issuesCmd.Flags().String("app", "", "Filter by GitHub App author")
	search_issuesCmd.Flags().Bool("archived", false, "Restrict search to archived repositories")
	search_issuesCmd.Flags().String("assignee", "", "Filter by assignee")
	search_issuesCmd.Flags().String("author", "", "Filter by author")
	search_issuesCmd.Flags().String("closed", "", "Filter on closed at `date`")
	search_issuesCmd.Flags().String("commenter", "", "Filter based on comments by `user`")
	search_issuesCmd.Flags().String("comments", "", "Filter on `number` of comments")
	search_issuesCmd.Flags().String("created", "", "Filter based on created at `date`")
	search_issuesCmd.Flags().Bool("include-prs", false, "Include pull requests in results")
	search_issuesCmd.Flags().String("interactions", "", "Filter on `number` of reactions and comments")
	search_issuesCmd.Flags().String("involves", "", "Filter based on involvement of `user`")
	search_issuesCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	search_issuesCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	search_issuesCmd.Flags().StringSlice("label", []string{}, "Filter on label")
	search_issuesCmd.Flags().String("language", "", "Filter based on the coding language")
	search_issuesCmd.Flags().IntP("limit", "L", 30, "Maximum number of results to fetch")
	search_issuesCmd.Flags().Bool("locked", false, "Filter on locked conversation status")
	search_issuesCmd.Flags().StringSlice("match", []string{}, "Restrict search to specific field of issue: {title|body|comments}")
	search_issuesCmd.Flags().String("mentions", "", "Filter based on `user` mentions")
	search_issuesCmd.Flags().String("milestone", "", "Filter by milestone `title`")
	search_issuesCmd.Flags().Bool("no-assignee", false, "Filter on missing assignee")
	search_issuesCmd.Flags().Bool("no-label", false, "Filter on missing label")
	search_issuesCmd.Flags().Bool("no-milestone", false, "Filter on missing milestone")
	search_issuesCmd.Flags().Bool("no-project", false, "Filter on missing project")
	search_issuesCmd.Flags().String("order", "desc", "Order of results returned, ignored unless '--sort' flag is specified: {asc|desc}")
	search_issuesCmd.Flags().String("owner", "", "Filter on repository owner")
	search_issuesCmd.Flags().String("project", "", "Filter on project board `number`")
	search_issuesCmd.Flags().String("reactions", "", "Filter on `number` of reactions")
	search_issuesCmd.Flags().StringSlice("repo", []string{}, "Filter on repository")
	search_issuesCmd.Flags().String("sort", "best-match", "Sort fetched results: {comments|created|interactions|reactions|reactions-+1|reactions--1|reactions-heart|reactions-smile|reactions-tada|reactions-thinking_face|updated}")
	search_issuesCmd.Flags().String("state", "", "Filter based on state: {open|closed}")
	search_issuesCmd.Flags().String("team-mentions", "", "Filter based on team mentions")
	search_issuesCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	search_issuesCmd.Flags().String("updated", "", "Filter on last updated at `date`")
	search_issuesCmd.Flags().StringSlice("visibility", []string{}, "Filter based on repository visibility: {public|private|internal}")
	search_issuesCmd.Flags().BoolP("web", "w", false, "Open the search query in the web browser")
	searchCmd.AddCommand(search_issuesCmd)

	// TODO app completion, project board number completion
	carapace.Gen(search_issuesCmd).FlagCompletion(carapace.ActionMap{
		"assignee": action.ActionSearchMultiRepo(search_issuesCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionAssignableUsers(cmd)
		}),
		"author":    gh.ActionUsers(gh.HostOpts{}),
		"closed":    action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"commenter": gh.ActionUsers(gh.HostOpts{}),
		"created":   action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"involves":  gh.ActionUsers(gh.HostOpts{}),
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSearchIssueFields().Invoke(c).Filter(c.Parts).ToA()
		}),
		"label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSearchMultiRepo(search_issuesCmd, func(cmd *cobra.Command) carapace.Action {
				return action.ActionLabels(cmd).Invoke(c).Filter(c.Parts).ToA()
			})
		}),
		"language": action.ActionLanguages(),
		"match": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("title", "body", "comments").Invoke(c).Filter(c.Parts).ToA()
		}),
		"mentions": gh.ActionUsers(gh.HostOpts{}),
		"milestone": action.ActionSearchMultiRepo(search_issuesCmd, func(cmd *cobra.Command) carapace.Action {
			return action.ActionMilestones(cmd)
		}),
		"order": carapace.ActionValues("asc", "desc"),
		"owner": gh.ActionOwners(gh.HostOpts{}),
		"repo": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			dummyCmd := &cobra.Command{}
			dummyCmd.Flags().String("repo", c.CallbackValue, "fake repo flag")
			return action.ActionOwnerRepositories(dummyCmd)
		}),
		"sort":    carapace.ActionValues("comments", "created", "interactions", "reactions", "reactions-+1", "reactions--1", "reactions-heart", "reactions-smile", "reactions-tada", "reactions-thinking_face", "updated"),
		"state":   carapace.ActionValues("open", "closed").StyleF(styles.Gh.ForState),
		"updated": action.ActionSearchRange(time.ActionDateTime(time.DateTimeOpts{Strict: true})),
		"visibility": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("public", "private", "internal").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
