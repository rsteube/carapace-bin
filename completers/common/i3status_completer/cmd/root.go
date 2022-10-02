package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "i3status",
	Short: "Generates a status line for i3bar, dzen2, xmobar or lemonbar",
	Long:  "https://i3wm.org/i3status/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("c", "c", "", "Specifies an alternate configuration file path.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionFiles(),
	})
}
