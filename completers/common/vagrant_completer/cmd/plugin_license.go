package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "install license",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_licenseCmd).Standalone()

	pluginCmd.AddCommand(plugin_licenseCmd)

	carapace.Gen(plugin_licenseCmd).PositionalCompletion(
		action.ActionPlugins(action.PluginOpts{Global: true}),
	)
}
