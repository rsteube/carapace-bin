package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Check keyboard and keymap for common mistakes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()

	lintCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	lintCmd.Flags().String("keyboard", "", "The keyboard to check.")
	lintCmd.Flags().String("keymap", "", "The keymap to check.")
	lintCmd.Flags().Bool("strict", false, "Treat warnings as errors.")
	rootCmd.AddCommand(lintCmd)

	carapace.Gen(lintCmd).FlagCompletion(carapace.ActionMap{
		"keyboard": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		}),
		"keymap": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeymaps(lintCmd.Flag("keyboard").Value.String())
		}),
	})
}
