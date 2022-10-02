package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Manage networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkCmd).Standalone()

	rootCmd.AddCommand(networkCmd)
}
