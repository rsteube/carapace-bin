package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Keyboard information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("ascii", false, "Render layout box drawings in ASCII only.")
	infoCmd.Flags().StringP("format", "f", "", "Format to display the data in")
	infoCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	infoCmd.Flags().String("keyboard", "", "Keyboard to show info for.")
	infoCmd.Flags().String("keymap", "", "Show the layers for a JSON keymap too.")
	infoCmd.Flags().BoolP("layouts", "l", false, "Render the layouts.")
	infoCmd.Flags().BoolP("matrix", "m", false, "Render the layouts with matrix information.")
	infoCmd.Flags().BoolP("rules-mk", "r", false, "Render the parsed values of the keyboard's rules.mk file.")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("friendly", "text", "json"),
		"keyboard": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		}),
		"keymap": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeymaps(compileCmd.Flag("keyboard").Value.String())
		}),
	})
}
