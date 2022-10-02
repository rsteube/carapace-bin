package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_editCmd).Standalone()
	pr_editCmd.Flags().StringSlice("add-assignee", []string{}, "Add assigned users by their `login`. Use \"@me\" to assign yourself.")
	pr_editCmd.Flags().StringSlice("add-label", []string{}, "Add labels by `name`")
	pr_editCmd.Flags().StringSlice("add-project", []string{}, "Add the pull request to projects by `name`")
	pr_editCmd.Flags().StringSlice("add-reviewer", []string{}, "Add reviewers by their `login`.")
	pr_editCmd.Flags().StringP("base", "B", "", "Change the base `branch` for this pull request")
	pr_editCmd.Flags().StringP("body", "b", "", "Set the new body.")
	pr_editCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	pr_editCmd.Flags().StringP("milestone", "m", "", "Edit the milestone the pull request belongs to by `name`")
	pr_editCmd.Flags().StringSlice("remove-assignee", []string{}, "Remove assigned users by their `login`. Use \"@me\" to unassign yourself.")
	pr_editCmd.Flags().StringSlice("remove-label", []string{}, "Remove labels by `name`")
	pr_editCmd.Flags().StringSlice("remove-project", []string{}, "Remove the pull request from projects by `name`")
	pr_editCmd.Flags().StringSlice("remove-reviewer", []string{}, "Remove reviewers by their `login`.")
	pr_editCmd.Flags().StringP("title", "t", "", "Set the new title.")
	prCmd.AddCommand(pr_editCmd)

	carapace.Gen(pr_editCmd).FlagCompletion(carapace.ActionMap{
		"add-assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAssignableUsers(pr_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"add-label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(pr_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"add-project": action.ActionProjects(pr_editCmd, action.ProjectOpts{Open: true}),
		"add-reviewer": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAssignableUsers(pr_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"base":      action.ActionBranches(pr_editCmd),
		"body":      action.ActionKeywordLinks(pr_editCmd),
		"body-file": carapace.ActionFiles(),
		"milestone": action.ActionMilestones(pr_editCmd),
		"remove-assignee": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionPullRequestAssignees(pr_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		"remove-label": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionPullRequestLabels(pr_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		"remove-reviewer": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionPullRequestReviewers(pr_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		// TODO remove-project
	})

	carapace.Gen(pr_editCmd).PositionalCompletion(
		action.ActionPullRequests(pr_editCmd, action.PullRequestOpts{Open: true}),
	)
}
