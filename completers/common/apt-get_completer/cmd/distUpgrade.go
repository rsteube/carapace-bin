package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var distUpgradeCmd = &cobra.Command{
	Use:   "dist-upgrade",
	Short: "dist-upgrade in addition to performing the function of upgrade",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distUpgradeCmd).Standalone()

	rootCmd.AddCommand(distUpgradeCmd)
}
