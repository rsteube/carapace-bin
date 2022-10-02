package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var remountCmd = &cobra.Command{
	Use:   "remount",
	Short: "remount partitions read-write",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remountCmd).Standalone()

	remountCmd.Flags().BoolS("R", "R", false, "automatically reboot the device if needed")

	rootCmd.AddCommand(remountCmd)
}
