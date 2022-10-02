package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Interact with the catalog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catalogCmd).Standalone()

	rootCmd.AddCommand(catalogCmd)
}
