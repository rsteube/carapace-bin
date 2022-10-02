package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Manage Consul's ACL policies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_policyCmd).Standalone()

	aclCmd.AddCommand(acl_policyCmd)
}
