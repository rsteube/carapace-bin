package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_rerunCmd = &cobra.Command{
	Use:   "rerun",
	Short: "Rerun a failed run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_rerunCmd).Standalone()
	run_rerunCmd.Flags().BoolP("debug", "d", false, "Rerun with debug logging")
	run_rerunCmd.Flags().Bool("failed", false, "Rerun only failed jobs, including dependencies")
	run_rerunCmd.Flags().StringP("job", "j", "", "Rerun a specific job from a run, including dependencies")
	runCmd.AddCommand(run_rerunCmd)

	carapace.Gen(run_rerunCmd).FlagCompletion(carapace.ActionMap{
		"job": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			return action.ActionWorkflowJobs(run_rerunCmd, c.Args[0], action.RunOpts{Failed: true, Successful: !run_rerunCmd.Flag("failed").Changed})
		}),
	})

	carapace.Gen(run_rerunCmd).PositionalCompletion(
		action.ActionWorkflowRuns(run_rerunCmd, action.RunOpts{Failed: true}),
	)
}
