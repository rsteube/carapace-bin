package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/qmk_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listLayoutsCmd = &cobra.Command{
	Use:   "list-keymaps",
	Short: "List the layouts for a specific keyboard",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listLayoutsCmd).Standalone()

	listLayoutsCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	listLayoutsCmd.Flags().String("keyboard", "", "Specify keyboard name.")
	rootCmd.AddCommand(listLayoutsCmd)

	carapace.Gen(listLayoutsCmd).FlagCompletion(carapace.ActionMap{
		"keyboard": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		}),
	})
}
