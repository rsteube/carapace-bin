package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var billingCmd = &cobra.Command{
	Use:   "billing",
	Short: "Manages the account payment methods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingCmd).Standalone()

	rootCmd.AddCommand(billingCmd)
}
