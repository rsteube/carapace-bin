package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var target_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed and available targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(target_listCmd).Standalone()

	target_listCmd.Flags().BoolP("help", "h", false, "Prints help information")
	target_listCmd.Flags().Bool("installed", false, "List only installed targets")
	target_listCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see")
	targetCmd.AddCommand(target_listCmd)

	carapace.Gen(target_listCmd).FlagCompletion(carapace.ActionMap{
		"toolchain": action.ActionToolchains(),
	})
}
