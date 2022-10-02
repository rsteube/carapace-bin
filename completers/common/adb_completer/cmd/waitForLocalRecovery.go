package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForLocalRecoveryCmd = &cobra.Command{
	Use:   "wait-for-local-recovery",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForLocalRecoveryCmd).Standalone()

	rootCmd.AddCommand(waitForLocalRecoveryCmd)
}
