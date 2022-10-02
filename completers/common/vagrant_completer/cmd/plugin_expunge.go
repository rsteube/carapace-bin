package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var plugin_expungeCmd = &cobra.Command{
	Use:   "expunge",
	Short: "expunge plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_expungeCmd).Standalone()

	plugin_expungeCmd.Flags().Bool("force", false, "Do not prompt for confirmation")
	plugin_expungeCmd.Flags().Bool("global-only", false, "Only expunge global plugins")
	plugin_expungeCmd.Flags().Bool("local", false, "Include plugins from local project for expunge")
	plugin_expungeCmd.Flags().Bool("local-only", false, "Only expunge local project plugins")
	plugin_expungeCmd.Flags().Bool("reinstall", false, "Reinstall current plugins after expunge")
	pluginCmd.AddCommand(plugin_expungeCmd)
}
