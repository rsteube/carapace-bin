package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "install the newest versions of all packages currently installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	rootCmd.AddCommand(upgradeCmd)
}
