package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var device_wifiCmd = &cobra.Command{
	Use:   "wifi",
	Short: "Perform operation on Wi-Fi devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_wifiCmd).Standalone()

	deviceCmd.AddCommand(device_wifiCmd)
}
