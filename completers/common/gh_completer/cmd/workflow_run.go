package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workflow_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a workflow by creating a workflow_dispatch event",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workflow_runCmd).Standalone()
	workflow_runCmd.Flags().StringArrayP("field", "F", []string{}, "Add a string parameter in `key=value` format, respecting @ syntax")
	workflow_runCmd.Flags().Bool("json", false, "Read workflow inputs as JSON via STDIN")
	workflow_runCmd.Flags().StringArrayP("raw-field", "f", []string{}, "Add a string parameter in `key=value` format")
	workflow_runCmd.Flags().StringP("ref", "r", "", "The branch or tag name which contains the version of the workflow file you'd like to run")
	workflowCmd.AddCommand(workflow_runCmd)

	carapace.Gen(workflow_runCmd).FlagCompletion(carapace.ActionMap{
		"ref": action.ActionBranches(workflow_runCmd),
	})

	carapace.Gen(workflow_runCmd).PositionalCompletion(
		action.ActionWorkflows(workflow_runCmd, action.WorkflowOpts{Enabled: true, Id: true, Name: true}),
	)
}
