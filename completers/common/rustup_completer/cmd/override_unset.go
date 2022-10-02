package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var override_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove the override toolchain for a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(override_unsetCmd).Standalone()

	override_unsetCmd.Flags().BoolP("help", "h", false, "Prints help information")
	override_unsetCmd.Flags().Bool("nonexistent", false, "Remove override toolchain for all nonexistent directories")
	override_unsetCmd.Flags().String("path", "", "Path to the directory")
	overrideCmd.AddCommand(override_unsetCmd)

	carapace.Gen(override_unsetCmd).FlagCompletion(carapace.ActionMap{
		"path": action.ActionOverrides(),
	})
}
