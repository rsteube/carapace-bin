package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var billing_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new credit card",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_addCmd).Standalone()

	billingCmd.AddCommand(billing_addCmd)
}
