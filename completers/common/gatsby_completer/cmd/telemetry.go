package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var telemetryCmd = &cobra.Command{
	Use:   "telemetry",
	Short: "Enable or disable Gatsby anonymous analytics collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telemetryCmd).Standalone()

	telemetryCmd.Flags().Bool("disable", false, "Disable telemetry")
	telemetryCmd.Flags().Bool("enable", false, "Enable telemetry")
	rootCmd.AddCommand(telemetryCmd)
}
