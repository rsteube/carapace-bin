package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var newKeyboardCmd = &cobra.Command{
	Use:   "new-keyboard",
	Short: "Creates a new keyboard directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newKeyboardCmd).Standalone()

	newKeyboardCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	newKeyboardCmd.Flags().String("keyboard", "", "Specify the name for the new keyboard directory")
	newKeyboardCmd.Flags().StringP("realname", "n", "", "Specify your real name if you want to use that.")
	newKeyboardCmd.Flags().StringP("type", "t", "", "Specify the keyboard type")
	newKeyboardCmd.Flags().StringP("username", "u", "", "Specify your username")
	rootCmd.AddCommand(newKeyboardCmd)

	carapace.Gen(newKeyboardCmd).FlagCompletion(carapace.ActionMap{
		"keyboard": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		}),
		"type": carapace.ActionValues("avr", "ps2avrgb"),
	})
}
