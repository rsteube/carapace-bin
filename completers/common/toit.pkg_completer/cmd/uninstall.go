package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/toit.pkg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls the package with the given name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalCompletion(
		action.ActionDependencies(uninstallCmd),
	)
}
