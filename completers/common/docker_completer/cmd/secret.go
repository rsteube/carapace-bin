package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Manage Docker secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretCmd).Standalone()

	rootCmd.AddCommand(secretCmd)
}
