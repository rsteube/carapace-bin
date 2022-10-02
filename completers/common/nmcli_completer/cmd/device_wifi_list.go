package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var device_wifi_wifi_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available Wi-Fi access points",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_wifi_wifi_listCmd).Standalone()

	device_wifi_wifi_listCmd.Flags().String("rescan", "", "trigger new Wi-Fi scan")
	device_wifiCmd.AddCommand(device_wifi_wifi_listCmd)

	carapace.Gen(device_wifi_wifi_listCmd).FlagCompletion(carapace.ActionMap{
		"rescan": carapace.ActionValues("yes", "no", "auto").StyleF(style.ForKeyword),
	})

	carapace.Gen(device_wifi_wifi_listCmd).PositionalCompletion(
		carapace.ActionValues("ifname", "bssid"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[0] {
			case "ifname":
				return net.ActionDevices(net.IncludedDevices{Wifi: true})
			case "bssid":
				return net.ActionBssids()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
