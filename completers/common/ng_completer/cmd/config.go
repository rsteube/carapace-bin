package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Retrieves or sets Angular configuration values in the angular.json file for the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().BoolP("global", "g", false, "Access the global configuration in the caller's home directory.")
	rootCmd.AddCommand(configCmd)
}
