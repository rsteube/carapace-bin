package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch a run until it completes, showing its progress",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_watchCmd).Standalone()
	run_watchCmd.Flags().Bool("exit-status", false, "Exit with non-zero status if run fails")
	run_watchCmd.Flags().IntP("interval", "i", 3, "Refresh interval in seconds")
	runCmd.AddCommand(run_watchCmd)

	carapace.Gen(run_watchCmd).PositionalCompletion(
		action.ActionWorkflowRuns(run_watchCmd, action.RunOpts{InProgress: true}),
	)
}
