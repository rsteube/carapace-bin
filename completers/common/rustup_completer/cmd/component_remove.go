package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var component_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a component from a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(component_removeCmd).Standalone()

	component_removeCmd.Flags().BoolP("help", "h", false, "Prints help information")
	component_removeCmd.Flags().String("target", "", "")
	component_removeCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see")
	componentCmd.AddCommand(component_removeCmd)

	carapace.Gen(component_removeCmd).FlagCompletion(carapace.ActionMap{
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionTargets(false).Invoke(c).ToMultiPartsA("-")
		}),
		"toolchain": action.ActionToolchains(),
	})

	carapace.Gen(component_removeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionComponents(true)
		}),
	)
}
