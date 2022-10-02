package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbacklight",
	Short: "adjust backlight brightness using RandR extension",
	Long:  "https://linux.die.net/man/1/xbacklight",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("d", "d", "", "Display to set backlight for.")
	rootCmd.Flags().StringS("dec", "dec", "", "Decreases brightness by the specified amount.")
	rootCmd.Flags().StringS("display", "display", "", "Display to set backlight for.")
	rootCmd.Flags().BoolS("get", "get", false, "Print  out  the  current backlight brightness of each output with such a control.")
	rootCmd.Flags().BoolS("help", "help", false, "Print out a summary of the usage and exit.")
	rootCmd.Flags().StringS("inc", "inc", "", "Increases brightness by the specified amount.")
	rootCmd.Flags().StringS("set", "set", "", "Sets each backlight brightness to the specified level.")
	rootCmd.Flags().StringS("steps", "steps", "", "Number of steps to take while fading.")
	rootCmd.Flags().StringS("time", "time", "", "Length of time to spend fading the backlight between old and new value.")
	rootCmd.Flags().BoolS("version", "version", false, "Print out the program version and exit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"d":       os.ActionDisplays(),
		"display": os.ActionDisplays(),
	})
}
