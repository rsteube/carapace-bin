package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_installCmd).Standalone()

	plugin_installCmd.Flags().String("alias", "", "Local name for plugin")
	plugin_installCmd.Flags().Bool("disable", false, "Do not enable the plugin on install")
	plugin_installCmd.Flags().Bool("disable-content-trust", false, "Skip image verification (default true)")
	plugin_installCmd.Flags().Bool("grant-all-permissions", false, "Grant all permissions necessary to run the plugin")
	pluginCmd.AddCommand(plugin_installCmd)

	carapace.Gen(plugin_installCmd).PositionalCompletion(
		docker.ActionPlugins(),
	)
}
