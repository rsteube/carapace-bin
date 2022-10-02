package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logTargetCmd = &cobra.Command{
	Use:   "log-target",
	Short: "Get/set logging target for manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logTargetCmd).Standalone()

	rootCmd.AddCommand(logTargetCmd)

	carapace.Gen(logTargetCmd).PositionalCompletion(
		action.ActionTargets(),
	)
}
