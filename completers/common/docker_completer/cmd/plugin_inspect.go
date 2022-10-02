package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display detailed information on one or more plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_inspectCmd).Standalone()

	plugin_inspectCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	pluginCmd.AddCommand(plugin_inspectCmd)

	carapace.Gen(plugin_inspectCmd).PositionalAnyCompletion(docker.ActionPlugins())
}
