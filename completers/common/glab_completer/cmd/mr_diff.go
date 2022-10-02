package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var mr_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "View changes in a merge request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_diffCmd).Standalone()
	mr_diffCmd.Flags().String("color", "auto", "Use color in diff output: {always|never|auto}")
	mrCmd.AddCommand(mr_diffCmd)

	carapace.Gen(mr_diffCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
	})

	carapace.Gen(mr_diffCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_diffCmd, ""),
	)
}
