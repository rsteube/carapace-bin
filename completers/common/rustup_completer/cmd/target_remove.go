package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var target_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a target from a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(target_removeCmd).Standalone()

	target_removeCmd.Flags().BoolP("help", "h", false, "Prints help information")
	target_removeCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see")
	targetCmd.AddCommand(target_removeCmd)

	carapace.Gen(target_removeCmd).FlagCompletion(carapace.ActionMap{
		"toolchain": action.ActionToolchains(),
	})

	carapace.Gen(target_removeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionTargets(true).Invoke(c).Filter(c.Args).ToMultiPartsA("-")
		}),
	)
}
