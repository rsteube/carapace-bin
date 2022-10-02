package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_disableCmd).Standalone()

	plugin_disableCmd.Flags().BoolP("force", "f", false, "Force the disable of an active plugin")
	pluginCmd.AddCommand(plugin_disableCmd)

	carapace.Gen(plugin_disableCmd).PositionalCompletion(
		docker.ActionPlugins(),
	)
}
