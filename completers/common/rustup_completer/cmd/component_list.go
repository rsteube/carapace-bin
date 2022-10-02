package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var component_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed and available components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(component_listCmd).Standalone()

	component_listCmd.Flags().BoolP("help", "h", false, "Prints help information")
	component_listCmd.Flags().Bool("installed", false, "List only installed components")
	component_listCmd.Flags().String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'. For more information see")
	componentCmd.AddCommand(component_listCmd)

	carapace.Gen(component_listCmd).FlagCompletion(carapace.ActionMap{
		"toolchain": action.ActionToolchains(),
	})
}
