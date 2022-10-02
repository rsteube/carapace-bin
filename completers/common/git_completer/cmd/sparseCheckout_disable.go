package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckout_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Repopulate the working directory with all files, disabling sparse checkouts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCheckout_disableCmd).Standalone()

	sparseCheckoutCmd.AddCommand(sparseCheckout_disableCmd)
}
