package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForBootloaderCmd = &cobra.Command{
	Use:   "wait-for-bootloader",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForBootloaderCmd).Standalone()

	rootCmd.AddCommand(waitForBootloaderCmd)
}
