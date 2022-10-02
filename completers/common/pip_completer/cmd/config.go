package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage local and global configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("global", false, "Use the system-wide configuration file only")
	configCmd.Flags().Bool("site", false, "Use the current environment configuration file only")
	configCmd.Flags().Bool("user", false, "Use the user configuration file only")
	rootCmd.AddCommand(configCmd)
}
