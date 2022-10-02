package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_createCmd).Standalone()
	issue_createCmd.Flags().StringSliceP("assignee", "a", []string{}, "Assign issue to people by their `usernames`")
	issue_createCmd.Flags().BoolP("confidential", "c", false, "Set an issue to be confidential. Default is false")
	issue_createCmd.Flags().StringP("description", "d", "", "Supply a description for issue")
	issue_createCmd.Flags().StringSliceP("label", "l", []string{}, "Add label by name. Multiple labels should be comma separated")
	issue_createCmd.Flags().String("link-type", "relates_to", "Type for the issue link")
	issue_createCmd.Flags().IntSlice("linked-issues", []int{}, "The IIDs of issues that this issue links to")
	issue_createCmd.Flags().Int("linked-mr", 0, "The IID of a merge request in which to resolve all issues")
	issue_createCmd.Flags().StringP("milestone", "m", "", "The global ID or title of a milestone to assign")
	issue_createCmd.Flags().Bool("no-editor", false, "Don't open editor to enter description. If set to true, uses prompt. Default is false")
	issue_createCmd.Flags().StringP("title", "t", "", "Supply a title for issue")
	issue_createCmd.Flags().Bool("web", false, "continue issue creation with web interface")
	issue_createCmd.Flags().IntP("weight", "w", 0, "The weight of the issue. Valid values are greater than or equal to 0.")
	issue_createCmd.Flags().BoolP("yes", "y", false, "Don't prompt for confirmation to submit the issue")
	issueCmd.AddCommand(issue_createCmd)

	carapace.Gen(issue_createCmd).FlagCompletion(carapace.ActionMap{
		// TODO more flags
		"assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjectMembers(issue_createCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(issue_createCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"link-type": carapace.ActionValues("relates_to", "blocks", "is_blocked_by"),
		"linked-issues": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionIssues(issue_createCmd, "opened").Invoke(c).Filter(c.Parts).ToA()
		}),
		"linked-mr": action.ActionMergeRequests(issue_createCmd, "opened"),
		"milestone": action.ActionMilestones(issue_createCmd),
	})
}
