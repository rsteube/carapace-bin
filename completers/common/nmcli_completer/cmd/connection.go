package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var connectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "NetworkManager's connections",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectionCmd).Standalone()

	rootCmd.AddCommand(connectionCmd)
}
