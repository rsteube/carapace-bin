package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverCmd).Standalone()

	rootCmd.AddCommand(serverCmd)
}
