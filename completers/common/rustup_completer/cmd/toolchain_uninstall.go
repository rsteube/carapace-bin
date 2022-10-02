package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var toolchain_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_uninstallCmd).Standalone()

	toolchain_uninstallCmd.Flags().BoolP("help", "h", false, "Prints help information")
	toolchainCmd.AddCommand(toolchain_uninstallCmd)

	carapace.Gen(toolchain_uninstallCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionToolchains().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
