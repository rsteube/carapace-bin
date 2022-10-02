package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade an existing plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_upgradeCmd).Standalone()

	plugin_upgradeCmd.Flags().Bool("disable-content-trust", false, "Skip image verification (default true)")
	plugin_upgradeCmd.Flags().Bool("grant-all-permissions", false, "Grant all permissions necessary to run the plugin")
	plugin_upgradeCmd.Flags().Bool("skip-remote-check", false, "Do not check if specified remote plugin matches existing plugin")
	pluginCmd.AddCommand(plugin_upgradeCmd)

	carapace.Gen(plugin_upgradeCmd).PositionalCompletion(
		docker.ActionPlugins(),
		carapace.ActionValues(),
	)
}
