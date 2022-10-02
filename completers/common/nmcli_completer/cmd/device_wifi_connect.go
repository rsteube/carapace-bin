package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_wifi_wifi_connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a Wi-Fi network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_wifi_wifi_connectCmd).Standalone()

	device_wifiCmd.AddCommand(device_wifi_wifi_connectCmd)

	carapace.Gen(device_wifi_wifi_connectCmd).PositionalCompletion(
		net.ActionSsids(),
	)
}
