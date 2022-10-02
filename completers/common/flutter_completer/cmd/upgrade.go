package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade your copy of Flutter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().BoolP("force", "f", false, "Force upgrade the flutter branch, potentially discarding local changes.")
	upgradeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	upgradeCmd.Flags().Bool("verify-only", false, "Checks for any new flutter updates, without actually fetching them.")
	rootCmd.AddCommand(upgradeCmd)
}
