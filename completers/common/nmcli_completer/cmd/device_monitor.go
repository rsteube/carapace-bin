package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor device activity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_monitorCmd).Standalone()

	deviceCmd.AddCommand(device_monitorCmd)

	carapace.Gen(device_monitorCmd).PositionalAnyCompletion(net.ActionDevices(net.AllDevices))
}
