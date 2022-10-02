package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var plugin_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install one or more plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_installCmd).Standalone()
	plugin_installCmd.PersistentFlags().Bool("exact", false, "Force installation of an exact version match (usually >= is accepted)")
	plugin_installCmd.PersistentFlags().StringP("file", "f", "", "Install a plugin from a tarball file, instead of downloading it")
	plugin_installCmd.PersistentFlags().Bool("reinstall", false, "Reinstall a plugin even if it already exists")
	plugin_installCmd.PersistentFlags().String("server", "", "A URL to download plugins from")
	pluginCmd.AddCommand(plugin_installCmd)

	carapace.Gen(plugin_installCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
