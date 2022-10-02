package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage the npm configuration files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()
	configCmd.PersistentFlags().BoolP("global", "g", false, "operate in global mode")

	rootCmd.AddCommand(configCmd)
}
