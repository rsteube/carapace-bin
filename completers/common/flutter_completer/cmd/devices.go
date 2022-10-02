package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List all connected devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicesCmd).Standalone()

	devicesCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach.")
	devicesCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	devicesCmd.Flags().Bool("machine", false, "Output device information in machine readable structured JSON format.")
	rootCmd.AddCommand(devicesCmd)
}
