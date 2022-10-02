package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var billing_setDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Make a credit card your default one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_setDefaultCmd).Standalone()

	billingCmd.AddCommand(billing_setDefaultCmd)

	// TODO positional completion
}
