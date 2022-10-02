package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconnect_deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "kick connection from device side to force reconnect",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconnect_deviceCmd).Standalone()

	reconnectCmd.AddCommand(reconnect_deviceCmd)
}
