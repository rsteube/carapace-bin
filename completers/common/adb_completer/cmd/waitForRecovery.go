package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForRecoveryCmd = &cobra.Command{
	Use:   "wait-for-recovery",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForRecoveryCmd).Standalone()

	rootCmd.AddCommand(waitForRecoveryCmd)
}
