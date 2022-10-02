package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall one or more Helm plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	pluginCmd.AddCommand(plugin_uninstallCmd)

	carapace.Gen(plugin_uninstallCmd).PositionalCompletion(
		action.ActionPlugins(),
	)
}
