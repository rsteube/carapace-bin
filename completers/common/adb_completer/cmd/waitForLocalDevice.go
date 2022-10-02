package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForLocalDeviceCmd = &cobra.Command{
	Use:   "wait-for-local-device",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForLocalDeviceCmd).Standalone()

	rootCmd.AddCommand(waitForLocalDeviceCmd)
}
