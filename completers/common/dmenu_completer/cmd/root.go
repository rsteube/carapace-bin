package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/dmenu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dmenu",
	Short: "dynamic menu",
	Long:  "https://linux.die.net/man/1/dmenu",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("b", "b", false, "dmenu appears at the bottom of the screen.")
	rootCmd.Flags().BoolS("f", "f", false, "dmenu grabs the keyboard before reading stdin if not reading from a tty.")
	rootCmd.Flags().StringS("fn", "fn", "", "defines the font or font set used.")
	rootCmd.Flags().StringS("h", "h", "", "dmenu uses a menu line of at least 'height' pixels tall, but no less than 8.")
	rootCmd.Flags().BoolS("i", "i", false, "dmenu matches menu items case insensitively.")
	rootCmd.Flags().StringS("l", "l", "", "dmenu lists items vertically, with the given number of lines.")
	rootCmd.Flags().StringS("m", "m", "", "dmenu is displayed on the monitor number supplied.")
	rootCmd.Flags().StringS("nb", "nb", "", "defines the normal background color.")
	rootCmd.Flags().StringS("nf", "nf", "", "defines the normal foreground color.")
	rootCmd.Flags().StringS("p", "p", "", "defines the prompt to be displayed to the left of the input field.")
	rootCmd.Flags().StringS("sb", "sb", "", "defines the selected background color.")
	rootCmd.Flags().StringS("sf", "sf", "", "defines the selected foreground color.")
	rootCmd.Flags().BoolS("v", "v", false, "prints version information to stdout, then exits.")
	rootCmd.Flags().StringS("w", "w", "", "sets the width of the dmenu window.")
	rootCmd.Flags().StringS("x", "x", "", "dmenu is placed at this offset measured from the left side of the monitor.")
	rootCmd.Flags().StringS("y", "y", "", "dmenu is placed at this offset measured from the top of the monitor.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"nb": action.ActionColors(),
		"nf": action.ActionColors(),
		"sb": action.ActionColors(),
		"sf": action.ActionColors(),
	})
}
