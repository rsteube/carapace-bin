package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var manCmd = &cobra.Command{
	Use:   "man",
	Short: "View the man page for a given command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manCmd).Standalone()

	manCmd.Flags().BoolP("help", "h", false, "Prints help information")
	manCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see")
	rootCmd.AddCommand(manCmd)

	carapace.Gen(manCmd).FlagCompletion(carapace.ActionMap{
		"toolchain": action.ActionToolchains(),
	})
}
