package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var plugin_repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "repair plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_repairCmd).Standalone()

	plugin_repairCmd.Flags().Bool("local", false, "Repair plugins in local project")
	pluginCmd.AddCommand(plugin_repairCmd)
}
