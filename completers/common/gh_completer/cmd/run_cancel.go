package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel a workflow run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_cancelCmd).Standalone()
	runCmd.AddCommand(run_cancelCmd)

	carapace.Gen(run_cancelCmd).PositionalCompletion(
		action.ActionWorkflowRuns(run_cancelCmd, action.RunOpts{InProgress: true}),
	)
}
