package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_reapplyCmd = &cobra.Command{
	Use:   "reapply",
	Short: "Attempt to update device with changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_reapplyCmd).Standalone()

	deviceCmd.AddCommand(device_reapplyCmd)

	carapace.Gen(device_reapplyCmd).PositionalCompletion(
		net.ActionDevices(net.AllDevices),
	)
}
