package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_enableCmd).Standalone()

	plugin_enableCmd.Flags().String("timeout", "", "HTTP client timeout (in seconds) (default 30)")
	pluginCmd.AddCommand(plugin_enableCmd)

	carapace.Gen(plugin_enableCmd).PositionalCompletion(
		docker.ActionPlugins(),
	)
}
