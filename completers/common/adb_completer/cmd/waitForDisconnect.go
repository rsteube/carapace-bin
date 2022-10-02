package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForDisconnectCmd = &cobra.Command{
	Use:   "wait-for-disconnect",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForDisconnectCmd).Standalone()

	rootCmd.AddCommand(waitForDisconnectCmd)
}
