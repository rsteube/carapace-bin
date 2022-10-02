package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/light_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "light",
	Short: "a program to control backlight controllers",
	Long:  "https://github.com/haikarainen/light",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("A", "A", "", "Increase brightness by value")
	rootCmd.Flags().BoolS("G", "G", false, "Get brightness")
	rootCmd.Flags().BoolS("I", "I", false, "Restore the previously saved brightness")
	rootCmd.Flags().BoolS("L", "L", false, "List available devices")
	rootCmd.Flags().BoolS("N", "N", false, "Set minimum brightness to value")
	rootCmd.Flags().BoolS("O", "O", false, "Save the current brightness")
	rootCmd.Flags().BoolS("P", "P", false, "Get minimum brightness")
	rootCmd.Flags().StringS("S", "S", "", "Set brightness to value")
	rootCmd.Flags().StringS("T", "T", "", "Multiply brightness by value")
	rootCmd.Flags().StringS("U", "U", "", "Decrease brightness by value")
	rootCmd.Flags().BoolS("V", "V", false, "Show program version and exit")
	rootCmd.Flags().BoolS("r", "r", false, "Interpret input and output values in raw mode")
	rootCmd.Flags().StringS("s", "s", "", "Specify device target path to use")
	rootCmd.Flags().StringS("v", "v", "", "Specify the verbosity level")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"s": action.ActionDevices(),
		"v": carapace.ActionValuesDescribed(
			"0", "Values only",
			"1", "Values, Errors",
			"2", "Values, Errors, Warnings",
			"3", "Values, Errors, Warnings, Notices",
		),
	})
}
