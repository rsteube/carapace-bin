package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var plugin_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_lsCmd).Standalone()
	plugin_lsCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	plugin_lsCmd.PersistentFlags().BoolP("project", "p", false, "List only the plugins used by the current project")
	pluginCmd.AddCommand(plugin_lsCmd)
}
