package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pass_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate a new password",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	generateCmd.Flags().BoolP("clip", "c", false, "additionally put on clipboard")
	generateCmd.Flags().BoolP("force", "f", false, "prompt before overwriting")
	generateCmd.Flags().BoolP("in-place", "i", false, "only replace first line of existing")
	generateCmd.Flags().BoolP("no-symbols", "n", false, "without symbols")
	rootCmd.AddCommand(generateCmd)

	carapace.Gen(generateCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPassNames().Invoke(c).ToMultiPartsA("/")
		}),
		carapace.ActionValues("8", "12", "16", "20", "24", "28", "32"),
	)
}
