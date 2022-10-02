package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var billing_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a credit card",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_rmCmd).Standalone()

	billingCmd.AddCommand(billing_rmCmd)

	// TODO positional completion
}
