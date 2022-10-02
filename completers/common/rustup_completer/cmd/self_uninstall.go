package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var self_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall rustup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_uninstallCmd).Standalone()

	self_uninstallCmd.Flags().BoolP("help", "h", false, "Prints help information")
	self_uninstallCmd.Flags().BoolS("y", "y", false, "")
	selfCmd.AddCommand(self_uninstallCmd)
}
