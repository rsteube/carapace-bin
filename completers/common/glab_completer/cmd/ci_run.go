package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Create or run a new CI pipeline",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_runCmd).Standalone()
	ci_runCmd.Flags().StringP("branch", "b", "", "Create pipeline on branch/ref <string>")
	ci_runCmd.Flags().StringSlice("variables", []string{}, "Pass variables to pipeline")
	ciCmd.AddCommand(ci_runCmd)

	carapace.Gen(ci_runCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_runCmd),
	})
}
