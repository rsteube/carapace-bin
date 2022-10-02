package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var trustCmd = &cobra.Command{
	Use:   "trust",
	Short: "Manage trust on Docker images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustCmd).Standalone()

	rootCmd.AddCommand(trustCmd)
}
