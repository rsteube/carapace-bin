package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconnectCmd = &cobra.Command{
	Use:   "reconnect",
	Short: "kick connection from host side to force reconnect",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconnectCmd).Standalone()

	rootCmd.AddCommand(reconnectCmd)
}
