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

	plugin_lsCmd.Flags().StringP("filter", "f", "", "Provide filter values (e.g. 'enabled=true')")
	plugin_lsCmd.Flags().String("format", "", "Pretty-print plugins using a Go template")
	plugin_lsCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	plugin_lsCmd.Flags().BoolP("quiet", "q", false, "Only display plugin IDs")
	pluginCmd.AddCommand(plugin_lsCmd)
}
