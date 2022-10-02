package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade all installed packages to the latest version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	rootCmd.AddCommand(upgradeCmd)
}
