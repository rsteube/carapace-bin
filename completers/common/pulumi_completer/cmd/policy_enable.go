package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var policy_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a Policy Pack for a Pulumi organization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_enableCmd).Standalone()
	policy_enableCmd.PersistentFlags().String("config", "", "The file path for the Policy Pack configuration file")
	policy_enableCmd.PersistentFlags().String("policy-group", "", "The Policy Group for which the Policy Pack will be enabled; if not specified, the default Policy Group is used")
	policyCmd.AddCommand(policy_enableCmd)
}
