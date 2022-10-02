package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var serviceLogTargetCmd = &cobra.Command{
	Use:   "service-log-target",
	Short: "Get/set logging target for service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceLogTargetCmd).Standalone()

	rootCmd.AddCommand(serviceLogTargetCmd)

	carapace.Gen(serviceLogTargetCmd).PositionalCompletion(
		action.ActionServices(),
		action.ActionTargets(),
	)
}
