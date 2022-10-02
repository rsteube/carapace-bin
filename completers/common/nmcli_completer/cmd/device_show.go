package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show details of device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_showCmd).Standalone()

	deviceCmd.AddCommand(device_showCmd)

	carapace.Gen(device_showCmd).PositionalCompletion(
		net.ActionDevices(net.AllDevices),
	)
}
