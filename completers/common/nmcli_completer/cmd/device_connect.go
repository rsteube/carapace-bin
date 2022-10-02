package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect the device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_connectCmd).Standalone()

	deviceCmd.AddCommand(device_connectCmd)

	carapace.Gen(device_connectCmd).PositionalCompletion(
		net.ActionDevices(net.AllDevices),
	)
}
