package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "View a running CI pipeline on current or other branch specified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_statusCmd).Standalone()
	ci_statusCmd.Flags().StringP("branch", "b", "", "Check pipeline status for a branch. (Default is current branch)")
	ci_statusCmd.Flags().BoolP("compact", "c", false, "Show status in compact format")
	ci_statusCmd.Flags().BoolP("live", "l", false, "Show status in real-time till pipeline ends")
	ciCmd.AddCommand(ci_statusCmd)

	carapace.Gen(ci_statusCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_statusCmd),
	})
}
