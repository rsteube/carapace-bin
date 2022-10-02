package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify one or more proties currently active on the device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_modifyCmd).Standalone()

	deviceCmd.AddCommand(device_modifyCmd)

	carapace.Gen(device_modifyCmd).PositionalCompletion(
		net.ActionDevices(net.AllDevices),
	)
}
