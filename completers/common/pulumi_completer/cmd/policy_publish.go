package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var policy_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a Policy Pack to the Pulumi service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_publishCmd).Standalone()
	policyCmd.AddCommand(policy_publishCmd)
}
