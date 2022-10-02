package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dselectUpgradeCmd = &cobra.Command{
	Use:   "dselect-upgrade",
	Short: "dselect-upgrade is used in conjunction with the traditional Debian packaging front-end",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dselectUpgradeCmd).Standalone()

	rootCmd.AddCommand(dselectUpgradeCmd)
}
