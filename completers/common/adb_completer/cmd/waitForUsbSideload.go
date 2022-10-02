package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForUsbSideloadCmd = &cobra.Command{
	Use:   "wait-for-usb-sideload",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForUsbSideloadCmd).Standalone()

	rootCmd.AddCommand(waitForUsbSideloadCmd)
}
