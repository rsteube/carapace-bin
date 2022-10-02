package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var override_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the override toolchain for a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(override_setCmd).Standalone()

	override_setCmd.Flags().BoolP("help", "h", false, "Prints help information")
	override_setCmd.Flags().String("path", "", "Path to the directory")
	overrideCmd.AddCommand(override_setCmd)

	carapace.Gen(override_setCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})

	carapace.Gen(override_setCmd).PositionalCompletion(
		action.ActionToolchains(),
	)
}
