package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Display which binary will be run for a given command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whichCmd).Standalone()

	whichCmd.Flags().BoolP("help", "h", false, "Prints help information")
	whichCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see")
	rootCmd.AddCommand(whichCmd)

	carapace.Gen(whichCmd).FlagCompletion(carapace.ActionMap{
		"toolchain": action.ActionToolchains(),
	})
}
