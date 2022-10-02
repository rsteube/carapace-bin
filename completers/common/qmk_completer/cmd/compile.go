package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile a QMK Firmware",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compileCmd).Standalone()

	compileCmd.Flags().BoolP("clean", "c", false, "Remove object files before compiling.")
	compileCmd.Flags().BoolP("dry-run", "n", false, "Don't actually build, just show the make command to be run.")
	compileCmd.Flags().StringP("env", "e", "", "Set a variable to be passed to make.")
	compileCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	compileCmd.Flags().String("keyboard", "", "The keyboard to build a firmware for.")
	compileCmd.Flags().String("keymap", "", "The keymap to build a firmware for.")
	compileCmd.Flags().StringP("parallel", "j", "", "Set the number of parallel make jobs.")
	rootCmd.AddCommand(compileCmd)

	carapace.Gen(compileCmd).FlagCompletion(carapace.ActionMap{
		"keyboard": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		}),
		"keymap": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeymaps(compileCmd.Flag("keyboard").Value.String())
		}),
	})

	carapace.Gen(compileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
