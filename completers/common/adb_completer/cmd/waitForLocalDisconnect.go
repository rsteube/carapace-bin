package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForLocalDisconnectCmd = &cobra.Command{
	Use:   "wait-for-local-disconnect",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForLocalDisconnectCmd).Standalone()

	rootCmd.AddCommand(waitForLocalDisconnectCmd)
}
