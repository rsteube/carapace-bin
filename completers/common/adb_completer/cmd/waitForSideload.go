package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForSideloadCmd = &cobra.Command{
	Use:   "wait-for-sideload",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForSideloadCmd).Standalone()

	rootCmd.AddCommand(waitForSideloadCmd)
}
