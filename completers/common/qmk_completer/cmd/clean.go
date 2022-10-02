package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean the QMK firmware folder of build artifacts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().BoolP("all", "a", false, "Remove *.hex and *.bin files in the QMK root as well.")
	cleanCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(cleanCmd)
}
