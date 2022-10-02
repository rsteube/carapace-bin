package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "list connected devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicesCmd).Standalone()

	devicesCmd.Flags().BoolS("l", "l", false, "long output")

	rootCmd.AddCommand(devicesCmd)
}
