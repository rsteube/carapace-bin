package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List project issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_listCmd).Standalone()
	issue_listCmd.Flags().BoolP("all", "A", false, "Get all issues")
	issue_listCmd.Flags().StringP("assignee", "a", "", "Filter issue by assignee <username>")
	issue_listCmd.Flags().String("author", "", "Filter issue by author <username>")
	issue_listCmd.Flags().BoolP("closed", "c", false, "Get only closed issues")
	issue_listCmd.Flags().BoolP("confidential", "C", false, "Filter by confidential issues")
	issue_listCmd.Flags().StringP("group", "g", "", "Get issues from group and it's subgroups")
	issue_listCmd.Flags().String("in", "title,description", "search in {title|description}")
	issue_listCmd.Flags().StringSliceP("label", "l", []string{}, "Filter issue by label <name>")
	issue_listCmd.Flags().StringP("milestone", "m", "", "Filter issue by milestone <id>")
	issue_listCmd.Flags().BoolP("mine", "M", false, "Filter only issues issues assigned to me")
	issue_listCmd.Flags().StringSlice("not-assignee", []string{}, "Filter issue by not being assigneed to <username>")
	issue_listCmd.Flags().StringSlice("not-author", []string{}, "Filter by not being by author(s) <username>")
	issue_listCmd.Flags().StringSlice("not-label", []string{}, "Filter issue by lack of label <name>")
	issue_listCmd.Flags().BoolP("opened", "o", false, "Get only opened issues")
	issue_listCmd.Flags().IntP("page", "p", 1, "Page number")
	issue_listCmd.Flags().IntP("per-page", "P", 30, "Number of items to list per page. (default 30)")
	issue_listCmd.Flags().String("search", "", "Search <string> in the fields defined by --in")
	issueCmd.AddCommand(issue_listCmd)

	carapace.Gen(issue_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee": action.ActionProjectMembers(issue_listCmd),
		"author":   action.ActionUsers(issue_listCmd),
		"group":    action.ActionGroups(issue_listCmd),
		"in": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("title", "description").Invoke(c).Filter(c.Parts).ToA()
		}),
		"label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"milestone": action.ActionMilestones(issue_listCmd),
		"not-assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjectMembers(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"not-author": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionUsers(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"not-label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(issue_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
