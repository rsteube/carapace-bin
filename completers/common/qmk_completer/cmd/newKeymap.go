package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var newKeymapCmd = &cobra.Command{
	Use:   "new-keymap",
	Short: "Creates a new keymap for the keyboard of your choosing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newKeymapCmd).Standalone()

	newKeymapCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	newKeymapCmd.Flags().String("keyboard", "", "Specify keyboard name.")
	newKeymapCmd.Flags().String("keymap", "", "Specify the name for the new keymap directory")
	rootCmd.AddCommand(newKeymapCmd)

	carapace.Gen(newKeymapCmd).FlagCompletion(carapace.ActionMap{
		"keyboard": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		}),
		"keymap": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeymaps(newKeymapCmd.Flag("keyboard").Value.String())
		}),
	})
}
