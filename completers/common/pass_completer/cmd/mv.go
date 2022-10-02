package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pass_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "rename or move old-path to new-path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mvCmd).Standalone()
	mvCmd.Flags().BoolP("force", "f", false, "mv forcefully")

	rootCmd.AddCommand(mvCmd)

	carapace.Gen(mvCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPassNames().Invoke(c).ToMultiPartsA("/")
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPassNames().Invoke(c).ToMultiPartsA("/")
		}),
	)
}
