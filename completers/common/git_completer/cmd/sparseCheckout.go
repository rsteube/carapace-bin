package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckoutCmd = &cobra.Command{
	Use:   "sparse-checkout",
	Short: "Initialize and modify the sparse-checkout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCheckoutCmd).Standalone()

	rootCmd.AddCommand(sparseCheckoutCmd)
}
