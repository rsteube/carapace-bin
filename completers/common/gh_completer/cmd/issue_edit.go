package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_editCmd).Standalone()
	issue_editCmd.Flags().StringSlice("add-assignee", []string{}, "Add assigned users by their `login`. Use \"@me\" to assign yourself.")
	issue_editCmd.Flags().StringSlice("add-label", []string{}, "Add labels by `name`")
	issue_editCmd.Flags().StringSlice("add-project", []string{}, "Add the issue to projects by `name`")
	issue_editCmd.Flags().StringP("body", "b", "", "Set the new body.")
	issue_editCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	issue_editCmd.Flags().StringP("milestone", "m", "", "Edit the milestone the issue belongs to by `name`")
	issue_editCmd.Flags().StringSlice("remove-assignee", []string{}, "Remove assigned users by their `login`. Use \"@me\" to unassign yourself.")
	issue_editCmd.Flags().StringSlice("remove-label", []string{}, "Remove labels by `name`")
	issue_editCmd.Flags().StringSlice("remove-project", []string{}, "Remove the issue from projects by `name`")
	issue_editCmd.Flags().StringP("title", "t", "", "Set the new title.")
	issueCmd.AddCommand(issue_editCmd)

	carapace.Gen(issue_editCmd).FlagCompletion(carapace.ActionMap{
		"add-assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAssignableUsers(issue_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"add-label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(issue_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"add-project": action.ActionProjects(issue_editCmd, action.ProjectOpts{Open: true}),
		"body":        action.ActionKeywordLinks(issue_editCmd),
		"body-file":   carapace.ActionFiles(),
		"milestone":   action.ActionMilestones(issue_editCmd),
		"remove-assignee": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionIssueAssignees(issue_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		"remove-label": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionIssueLabels(issue_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		// TODO remove-project
	})

	carapace.Gen(issue_editCmd).PositionalCompletion(
		action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true}),
	)
}
