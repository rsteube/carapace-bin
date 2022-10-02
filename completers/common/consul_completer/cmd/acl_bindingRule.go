package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_bindingRuleCmd = &cobra.Command{
	Use:   "binding-rule",
	Short: "Manage Consul's ACL binding rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_bindingRuleCmd).Standalone()

	aclCmd.AddCommand(acl_bindingRuleCmd)
}
