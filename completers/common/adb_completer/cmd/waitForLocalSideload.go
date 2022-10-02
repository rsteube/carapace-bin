package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitForLocalSideloadCmd = &cobra.Command{
	Use:   "wait-for-local-sideload",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForLocalSideloadCmd).Standalone()

	rootCmd.AddCommand(waitForLocalSideloadCmd)
}
