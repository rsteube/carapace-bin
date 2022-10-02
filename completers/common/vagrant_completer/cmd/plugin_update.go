package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_updateCmd).Standalone()

	plugin_updateCmd.Flags().Bool("local", false, "Update plugin in local project")
	pluginCmd.AddCommand(plugin_updateCmd)

	carapace.Gen(plugin_updateCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPlugins(action.PluginOpts{
				Local:  plugin_updateCmd.Flag("local").Changed,
				Global: !plugin_updateCmd.Flag("local").Changed,
			})
		}),
	)
}
