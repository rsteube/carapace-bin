package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "reboot the device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebootCmd).Standalone()

	rootCmd.AddCommand(rebootCmd)

	carapace.Gen(rebootCmd).PositionalCompletion(
		carapace.ActionValues("bootloader", "recovery", "sideload", "sideload-auto-reboot"),
	)
}
