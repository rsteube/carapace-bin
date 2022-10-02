package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_updateCmd).Standalone()
	issue_updateCmd.Flags().StringSliceP("assignee", "a", []string{}, "assign users via username, prefix with '!' or '-' to remove from existing assignees, '+' to add, otherwise replace existing assignees with given users")
	issue_updateCmd.Flags().BoolP("confidential", "c", false, "Make issue confidential")
	issue_updateCmd.Flags().StringP("description", "d", "", "Issue description")
	issue_updateCmd.Flags().StringSliceP("label", "l", []string{}, "add labels")
	issue_updateCmd.Flags().Bool("lock-discussion", false, "Lock discussion on issue")
	issue_updateCmd.Flags().StringP("milestone", "m", "", "title of the milestone to assign, pass \"\" or 0 to unassign")
	issue_updateCmd.Flags().BoolP("public", "p", false, "Make issue public")
	issue_updateCmd.Flags().StringP("title", "t", "", "Title of issue")
	issue_updateCmd.Flags().Bool("unassign", false, "unassign all users")
	issue_updateCmd.Flags().StringSliceP("unlabel", "u", []string{}, "remove labels")
	issue_updateCmd.Flags().Bool("unlock-discussion", false, "Unlock discussion on issue")
	issueCmd.AddCommand(issue_updateCmd)

	carapace.Gen(issue_updateCmd).FlagCompletion(carapace.ActionMap{
		"assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjectMembers(issue_updateCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(issue_updateCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"milestone": action.ActionMilestones(issue_updateCmd),
		"unlabel": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(issue_updateCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(issue_updateCmd).PositionalCompletion(
		action.ActionIssues(issue_updateCmd, "opened"),
	)
}
