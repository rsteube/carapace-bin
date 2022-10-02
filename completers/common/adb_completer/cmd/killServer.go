package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var killServerCmd = &cobra.Command{
	Use:   "kill-server",
	Short: "kill the server if it is running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killServerCmd).Standalone()

	rootCmd.AddCommand(killServerCmd)
}
