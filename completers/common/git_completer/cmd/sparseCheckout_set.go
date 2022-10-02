package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckout_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Enable the necessary sparse-checkout config settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCheckout_setCmd).Standalone()

	sparseCheckoutCmd.AddCommand(sparseCheckout_setCmd)

	carapace.Gen(sparseCheckout_setCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
