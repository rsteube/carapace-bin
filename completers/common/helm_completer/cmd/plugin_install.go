package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var plugin_installCmd = &cobra.Command{
	Use:   "install",
	Short: "install one or more Helm plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	plugin_installCmd.Flags().String("version", "", "specify a version constraint. If this is not specified, the latest version is installed")
	pluginCmd.AddCommand(plugin_installCmd)

	carapace.Gen(plugin_installCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
