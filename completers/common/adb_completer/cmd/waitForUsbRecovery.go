package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForUsbRecoveryCmd = &cobra.Command{
	Use:   "wait-for-usb-recovery",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForUsbRecoveryCmd).Standalone()

	rootCmd.AddCommand(waitForUsbRecoveryCmd)
}
