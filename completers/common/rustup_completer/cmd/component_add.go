package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var component_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a component to a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(component_addCmd).Standalone()

	component_addCmd.Flags().BoolP("help", "h", false, "Prints help information")
	component_addCmd.Flags().String("target", "", "")
	component_addCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see")
	componentCmd.AddCommand(component_addCmd)

	carapace.Gen(component_addCmd).FlagCompletion(carapace.ActionMap{
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionTargets(false).Invoke(c).ToMultiPartsA("-")
		}),
		"toolchain": action.ActionToolchains(),
	})

	carapace.Gen(component_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionAvailableComponents().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
