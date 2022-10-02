package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set and get glab settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()
	configCmd.Flags().BoolP("global", "g", false, "use global config file")
	rootCmd.AddCommand(configCmd)
}
