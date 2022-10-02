package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace a CI job log in real time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_traceCmd).Standalone()
	ci_traceCmd.Flags().StringP("branch", "b", "", "Check pipeline status for a branch. (Default is the current branch)")
	ciCmd.AddCommand(ci_traceCmd)

	carapace.Gen(ci_statusCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_statusCmd),
	})
}
