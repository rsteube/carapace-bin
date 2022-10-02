package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var plugin_installCmd = &cobra.Command{
	Use:   "install",
	Short: "install plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_installCmd).Standalone()

	plugin_installCmd.Flags().String("entry-point", "", "The name of the entry point file for loading the plugin.")
	plugin_installCmd.Flags().Bool("local", false, "Install plugin for local project only")
	plugin_installCmd.Flags().Bool("plugin-clean-sources", false, "Remove all plugin sources defined so far (including defaults)")
	plugin_installCmd.Flags().String("plugin-source", "", "Add a RubyGems repository source")
	plugin_installCmd.Flags().String("plugin-version", "", "Install a specific version of the plugin")
	plugin_installCmd.Flags().Bool("verbose", false, "Enable verbose output for plugin installation")
	pluginCmd.AddCommand(plugin_installCmd)
}
