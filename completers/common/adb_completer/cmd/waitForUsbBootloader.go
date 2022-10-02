package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForUsbBootloaderCmd = &cobra.Command{
	Use:   "wait-for-usb-bootloader",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForUsbBootloaderCmd).Standalone()

	rootCmd.AddCommand(waitForUsbBootloaderCmd)
}
