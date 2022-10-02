package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Server access configurations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(contextCmd).Standalone()

	rootCmd.AddCommand(contextCmd)
}
