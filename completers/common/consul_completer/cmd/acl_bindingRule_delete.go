package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_bindingRule_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an ACL binding rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_bindingRule_deleteCmd).Standalone()
	addClientFlags(acl_bindingRule_deleteCmd)
	addServerFlags(acl_bindingRule_deleteCmd)

	acl_bindingRule_deleteCmd.Flags().String("id", "", "The ID of the binding rule to delete.")
	acl_bindingRule_deleteCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_bindingRuleCmd.AddCommand(acl_bindingRule_deleteCmd)
}
