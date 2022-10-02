package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForUsbRescueCmd = &cobra.Command{
	Use:   "wait-for-usb-rescue",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForUsbRescueCmd).Standalone()

	rootCmd.AddCommand(waitForUsbRescueCmd)
}
