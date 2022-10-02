package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listKeymapsCmd = &cobra.Command{
	Use:   "list-keymaps",
	Short: "List the keymaps for a specific keyboard",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listKeymapsCmd).Standalone()

	listKeymapsCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	listKeymapsCmd.Flags().String("keyboard", "", "Specify keyboard name.")
	rootCmd.AddCommand(listKeymapsCmd)

	carapace.Gen(listKeymapsCmd).FlagCompletion(carapace.ActionMap{
		"keyboard": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		}),
	})
}
