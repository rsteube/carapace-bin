package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_forCmd = &cobra.Command{
	Use:   "for",
	Short: "Create new merge request for an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_forCmd).Standalone()
	mr_forCmd.Flags().Bool("allow-collaboration", false, "Allow commits from other members")
	mr_forCmd.Flags().StringP("assignee", "a", "", "Assign merge request to people by their IDs. Multiple values should be comma separated ")
	mr_forCmd.Flags().Bool("draft", true, "Mark merge request as a draft. Default is true")
	mr_forCmd.Flags().StringP("label", "l", "", "Add label by name. Multiple labels should be comma separated")
	mr_forCmd.Flags().IntP("milestone", "m", -1, "add milestone by <id> for merge request")
	mr_forCmd.Flags().Bool("remove-source-branch", false, "Remove Source Branch on merge")
	mr_forCmd.Flags().StringP("target-branch", "b", "", "The target or base branch into which you want your code merged")
	mr_forCmd.Flags().Bool("wip", false, "Mark merge request as a work in progress. Overrides --draft")
	mr_forCmd.Flags().Bool("with-labels", false, "Copy labels from issue to the merge request")
	mrCmd.AddCommand(mr_forCmd)

	carapace.Gen(mr_forCmd).FlagCompletion(carapace.ActionMap{
		"assignee": action.ActionProjectMembers(mr_forCmd),
		"label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(mr_forCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"target-branch": action.ActionBranches(mr_forCmd),
	})

	carapace.Gen(mr_forCmd).PositionalCompletion(
		action.ActionIssues(mr_forCmd, "opened"),
	)
}
