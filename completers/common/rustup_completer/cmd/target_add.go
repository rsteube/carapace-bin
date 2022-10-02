package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var target_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a target to a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(target_addCmd).Standalone()

	target_addCmd.Flags().BoolP("help", "h", false, "Prints help information")
	target_addCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see")
	targetCmd.AddCommand(target_addCmd)

	carapace.Gen(target_addCmd).FlagCompletion(carapace.ActionMap{
		"toolchain": action.ActionToolchains(),
	})

	carapace.Gen(target_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionTargets(false).Invoke(c).Filter(c.Args).ToMultiPartsA("-")
		}),
	)
}
