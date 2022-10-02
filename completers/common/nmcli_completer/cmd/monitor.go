package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor NetworkManager changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitorCmd).Standalone()

	rootCmd.AddCommand(monitorCmd)
}
