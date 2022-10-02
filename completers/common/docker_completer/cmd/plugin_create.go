package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a plugin from a rootfs and configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_createCmd).Standalone()

	plugin_createCmd.Flags().Bool("compress", false, "Compress the context using gzip")
	pluginCmd.AddCommand(plugin_createCmd)

	carapace.Gen(plugin_createCmd).PositionalCompletion(
		docker.ActionPlugins(),
		carapace.ActionDirectories(),
	)
}
