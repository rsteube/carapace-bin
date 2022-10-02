package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workflow_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the summary of a workflow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workflow_viewCmd).Standalone()
	workflow_viewCmd.Flags().StringP("ref", "r", "", "The branch or tag name which contains the version of the workflow file you'd like to view")
	workflow_viewCmd.Flags().BoolP("web", "w", false, "Open workflow in the browser")
	workflow_viewCmd.Flags().BoolP("yaml", "y", false, "View the workflow yaml file")
	workflowCmd.AddCommand(workflow_viewCmd)

	carapace.Gen(workflow_viewCmd).FlagCompletion(carapace.ActionMap{
		"ref": action.ActionBranches(workflow_viewCmd),
	})

	carapace.Gen(workflow_viewCmd).PositionalCompletion(
		action.ActionWorkflows(workflow_viewCmd, action.WorkflowOpts{Enabled: true, Id: true, Name: true}),
	)
}
