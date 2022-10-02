package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_policy_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an ACL policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_policy_deleteCmd).Standalone()
	addClientFlags(acl_policy_deleteCmd)
	addServerFlags(acl_policy_deleteCmd)

	acl_policy_deleteCmd.Flags().String("id", "", "The ID of the policy to delete.")
	acl_policy_deleteCmd.Flags().String("name", "", "The name of the policy to delete.")
	acl_policy_deleteCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_policyCmd.AddCommand(acl_policy_deleteCmd)
}
