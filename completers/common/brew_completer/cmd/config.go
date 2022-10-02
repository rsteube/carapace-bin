package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show Homebrew and system configuration info useful for debugging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	configCmd.Flags().BoolP("help", "h", false, "Show this message.")
	configCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	configCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(configCmd)
}
