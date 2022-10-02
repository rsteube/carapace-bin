package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconnect_offlineCmd = &cobra.Command{
	Use:   "offline",
	Short: "reset offline/unauthorized devices to force reconnect",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconnect_offlineCmd).Standalone()

	reconnectCmd.AddCommand(reconnect_offlineCmd)
}
