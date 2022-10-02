package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var billing_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show all of your credit cards",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_lsCmd).Standalone()

	billingCmd.AddCommand(billing_lsCmd)
}
