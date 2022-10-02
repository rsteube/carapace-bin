package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var self_upgradeDataCmd = &cobra.Command{
	Use:   "upgrade-data",
	Short: "Upgrade the internal data format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_upgradeDataCmd).Standalone()

	self_upgradeDataCmd.Flags().BoolP("help", "h", false, "Prints help information")
	selfCmd.AddCommand(self_upgradeDataCmd)
}
