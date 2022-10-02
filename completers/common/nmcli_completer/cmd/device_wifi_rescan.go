package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_wifi_rescanCmd = &cobra.Command{
	Use:   "rescan",
	Short: "Request that NetworkManager immediately re-scan for available access points",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_wifi_rescanCmd).Standalone()

	device_wifiCmd.AddCommand(device_wifi_rescanCmd)

	carapace.Gen(device_wifi_rescanCmd).PositionalCompletion(
		carapace.ActionValues("ifname"),
		net.ActionDevices(net.IncludedDevices{Wifi: true}),
	)

	carapace.Gen(device_wifi_rescanCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args)%2 == 0 {
				return carapace.ActionValues("ssid")
			} else {
				return net.ActionSsids()
			}
		}),
	)
}
