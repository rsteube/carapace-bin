package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operator_autopilotCmd = &cobra.Command{
	Use:   "autopilot",
	Short: "Provides tools for modifying Autopilot configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operator_autopilotCmd).Standalone()

	operatorCmd.AddCommand(operator_autopilotCmd)
}
