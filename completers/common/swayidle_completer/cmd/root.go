package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swayidle",
	Short: "Idle manager for Wayland",
	Long:  "https://github.com/swaywm/swayidle",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("C", "C", "", "path to config file")
	rootCmd.Flags().StringS("S", "S", "", "pick the seat to work with")
	rootCmd.Flags().BoolS("d", "d", false, "debug")
	rootCmd.Flags().BoolS("h", "h", false, "this help menu")
	rootCmd.Flags().BoolS("w", "w", false, "wait for command to finish")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionFiles(),
	})
}
