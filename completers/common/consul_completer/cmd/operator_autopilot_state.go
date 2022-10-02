package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operator_autopilot_stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Display the current Autopilot configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operator_autopilot_stateCmd).Standalone()
	addClientFlags(operator_autopilot_stateCmd)
	addServerFlags(operator_autopilot_stateCmd)
	operator_autopilot_stateCmd.Flags().String("format", "", "Output format")

	operator_autopilotCmd.AddCommand(operator_autopilot_stateCmd)

	carapace.Gen(operator_autopilot_stateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
