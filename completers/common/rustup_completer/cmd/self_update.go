package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var self_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Download and install updates to rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_updateCmd).Standalone()

	self_updateCmd.Flags().BoolP("help", "h", false, "Prints help information")
	selfCmd.AddCommand(self_updateCmd)
}
