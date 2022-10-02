package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var c2jsonCmd = &cobra.Command{
	Use:   "c2json",
	Short: "Creates a keymap.json from a keymap.c file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(c2jsonCmd).Standalone()

	c2jsonCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	c2jsonCmd.Flags().String("keyboard", "", "The keyboard's name")
	c2jsonCmd.Flags().String("keymap", "", "The keymap's name")
	c2jsonCmd.Flags().Bool("no-cpp", false, "Do not use 'cpp' on keymap.c")
	c2jsonCmd.Flags().StringP("output", "o", "", "File to write to")
	c2jsonCmd.Flags().BoolP("quiet", "q", false, "Quiet mode, only output error messages")
	rootCmd.AddCommand(c2jsonCmd)

	carapace.Gen(c2jsonCmd).FlagCompletion(carapace.ActionMap{
		"keyboard": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		}),
		"keymap": action.ActionKeymaps(""),
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(c2jsonCmd).PositionalCompletion(
		carapace.ActionFiles(".c"),
	)
}
