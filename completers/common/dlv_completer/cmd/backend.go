package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Help about the --backend flag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backendCmd).Standalone()
	rootCmd.AddCommand(backendCmd)
}
