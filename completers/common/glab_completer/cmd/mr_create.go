package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new merge request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_createCmd).Standalone()
	mr_createCmd.Flags().Bool("allow-collaboration", false, "Allow commits from other members")
	mr_createCmd.Flags().StringSliceP("assignee", "a", []string{}, "Assign merge request to people by their `usernames`")
	mr_createCmd.Flags().Bool("copy-issue-labels", false, "Copy labels from issue to the merge request. Used with --related-issue")
	mr_createCmd.Flags().Bool("create-source-branch", false, "Create source branch if it does not exist")
	mr_createCmd.Flags().StringP("description", "d", "", "Supply a description for merge request")
	mr_createCmd.Flags().Bool("draft", false, "Mark merge request as a draft")
	mr_createCmd.Flags().BoolP("fill", "f", false, "Do not prompt for title/description and just use commit info")
	mr_createCmd.Flags().Bool("fill-commit-body", false, "Fill description with each commit body when multiple commits. Can only be used with --fill")
	mr_createCmd.Flags().StringP("head", "H", "", "Select another head repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or the project ID or full URL")
	mr_createCmd.Flags().StringSliceP("label", "l", []string{}, "Add label by name. Multiple labels should be comma separated")
	mr_createCmd.Flags().StringP("milestone", "m", "", "The global ID or title of a milestone to assign")
	mr_createCmd.Flags().Bool("no-editor", false, "Don't open editor to enter description. If set to true, uses prompt. Default is false")
	mr_createCmd.Flags().Bool("push", false, "Push committed changes after creating merge request. Make sure you have committed changes")
	mr_createCmd.Flags().StringP("related-issue", "i", "", "Create merge request for an issue. The merge request title will be created from the issue if --title is not provided.")
	mr_createCmd.Flags().Bool("remove-source-branch", false, "Remove Source Branch on merge")
	mr_createCmd.Flags().StringSlice("reviewer", []string{}, "Request review from users by their `usernames`")
	mr_createCmd.Flags().StringP("source-branch", "s", "", "The Branch you are creating the merge request. Default is the current branch.")
	mr_createCmd.Flags().Bool("squash-before-merge", false, "Squash commits into a single commit when merging")
	mr_createCmd.Flags().StringP("target-branch", "b", "", "The target or base branch into which you want your code merged")
	mr_createCmd.Flags().String("target-project", "", "Add target project by id or OWNER/REPO or GROUP/NAMESPACE/REPO")
	mr_createCmd.Flags().StringP("title", "t", "", "Supply a title for merge request")
	mr_createCmd.Flags().BoolP("web", "w", false, "continue merge request creation on web browser")
	mr_createCmd.Flags().Bool("wip", false, "Mark merge request as a work in progress. Alternative to --draft")
	mr_createCmd.Flags().BoolP("yes", "y", false, "Skip submission confirmation prompt, with --fill it skips all optional prompts")
	mrCmd.AddCommand(mr_createCmd)

	carapace.Gen(mr_createCmd).FlagCompletion(carapace.ActionMap{
		"assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjectMembers(mr_createCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(mr_createCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"milestone":     action.ActionMilestones(mr_createCmd),
		"related-issue": action.ActionIssues(mr_createCmd, "opened"),
		"reviewer": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProjectMembers(mr_createCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"source-branch": action.ActionBranches(mr_createCmd),
		"target-branch": action.ActionBranches(mr_createCmd),
	})
}
