package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceCmd).Standalone()

	rootCmd.AddCommand(serviceCmd)
}
