package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pass_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "copies old-path to new-path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()
	cpCmd.Flags().BoolP("force", "f", false, "cp forcefully")

	rootCmd.AddCommand(cpCmd)

	carapace.Gen(cpCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPassNames().Invoke(c).ToMultiPartsA("/")
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPassNames().Invoke(c).ToMultiPartsA("/")
		}),
	)
}
