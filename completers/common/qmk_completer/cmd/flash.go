package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var flashCmd = &cobra.Command{
	Use:   "flash",
	Short: "QMK Flash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flashCmd).Standalone()

	flashCmd.Flags().String("bootloader", "", "The flash command, corresponding to qmk's make options of bootloaders.")
	flashCmd.Flags().BoolP("bootloaders", "b", false, "List the available bootloaders.")
	flashCmd.Flags().BoolP("clean", "c", false, "Remove object files before compiling.")
	flashCmd.Flags().BoolP("dry-run", "n", false, "Don't actually build")
	flashCmd.Flags().StringP("env", "e", "", "Set a variable to be passed to make.")
	flashCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	flashCmd.Flags().String("keyboard", "", "The keyboard to build a firmware for")
	flashCmd.Flags().String("keymap", "", "The keymap to build a firmware for")
	flashCmd.Flags().StringP("parallel", "j", "", "Set the number of parallel make jobs")
	rootCmd.AddCommand(flashCmd)

	carapace.Gen(flashCmd).FlagCompletion(carapace.ActionMap{
		"bootloader": action.ActionBootloaders(),
		"keyboard": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		}),
		"keymap": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeymaps(flashCmd.Flag("keyboard").Value.String())
		}),
	})

	carapace.Gen(flashCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
