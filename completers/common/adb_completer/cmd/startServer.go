package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var startServerCmd = &cobra.Command{
	Use:   "start-server",
	Short: "ensure that there is a server running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startServerCmd).Standalone()

	rootCmd.AddCommand(startServerCmd)
}
