package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForRescueCmd = &cobra.Command{
	Use:   "wait-for-rescue",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForRescueCmd).Standalone()

	rootCmd.AddCommand(waitForRescueCmd)
}
