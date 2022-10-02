package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var policy_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a Policy Pack for a Pulumi organization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_disableCmd).Standalone()
	policy_disableCmd.PersistentFlags().String("policy-group", "", "The Policy Group for which the Policy Pack will be disabled; if not specified, the default Policy Group is used")
	policy_disableCmd.PersistentFlags().String("version", "", "The version of the Policy Pack that will be disabled; if not specified, any enabled version of the Policy Pack will be disabled")
	policyCmd.AddCommand(policy_disableCmd)
}
