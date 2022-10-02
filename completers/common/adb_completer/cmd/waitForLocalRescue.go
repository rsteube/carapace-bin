package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForLocalRescueCmd = &cobra.Command{
	Use:   "wait-for-local-rescue",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForLocalRescueCmd).Standalone()

	rootCmd.AddCommand(waitForLocalRescueCmd)
}
