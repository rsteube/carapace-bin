package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForUsbDisconnectCmd = &cobra.Command{
	Use:   "wait-for-usb-disconnect",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForUsbDisconnectCmd).Standalone()

	rootCmd.AddCommand(waitForUsbDisconnectCmd)
}
