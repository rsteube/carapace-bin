package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operator_autopilot_getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "Display the current Autopilot configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operator_autopilot_getConfigCmd).Standalone()
	addClientFlags(operator_autopilot_getConfigCmd)
	addServerFlags(operator_autopilot_getConfigCmd)

	operator_autopilotCmd.AddCommand(operator_autopilot_getConfigCmd)
}
