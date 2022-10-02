package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sparseCheckout_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Describe the directories or patterns in the sparse-checkout file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCheckout_listCmd).Standalone()

	sparseCheckoutCmd.AddCommand(sparseCheckout_listCmd)
}
