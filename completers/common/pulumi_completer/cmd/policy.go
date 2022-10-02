package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Manage resource policies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policyCmd).Standalone()
	rootCmd.AddCommand(policyCmd)
}
