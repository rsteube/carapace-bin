package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckout_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add all files under SOME/DIR/ECTORY/ (at any depth) to the sparse checkout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCheckout_addCmd).Standalone()

	sparseCheckoutCmd.AddCommand(sparseCheckout_addCmd)

	carapace.Gen(sparseCheckout_addCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
