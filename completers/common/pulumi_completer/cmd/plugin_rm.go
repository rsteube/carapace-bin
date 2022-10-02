package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more plugins from the download cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_rmCmd).Standalone()
	plugin_rmCmd.PersistentFlags().BoolP("all", "a", false, "Remove all plugins")
	plugin_rmCmd.PersistentFlags().BoolP("yes", "y", false, "Skip confirmation prompts, and proceed with removal anyway")
	pluginCmd.AddCommand(plugin_rmCmd)

	carapace.Gen(plugin_rmCmd).PositionalCompletion(
		action.ActionPluginKinds(plugin_rmCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPluginNames(plugin_rmCmd, c.Args[0])
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPluginVersions(plugin_rmCmd, c.Args[0], c.Args[1])
		}),
	)
}
