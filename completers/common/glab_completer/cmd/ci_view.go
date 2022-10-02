package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View, run, trace/logs, and cancel CI jobs current pipeline",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_viewCmd).Standalone()
	ci_viewCmd.Flags().StringP("branch", "b", "", "Check pipeline status for a branch/tag. (Default is the current branch)")
	ciCmd.AddCommand(ci_viewCmd)

	carapace.Gen(ci_viewCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_viewCmd),
	})

	carapace.Gen(ci_viewCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionBranches(ci_viewCmd),
			action.ActionReleases(ci_viewCmd),
		).ToA(),
	)
}
