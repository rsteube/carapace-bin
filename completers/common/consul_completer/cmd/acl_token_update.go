package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var acl_token_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an ACL token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_token_updateCmd).Standalone()
	addClientFlags(acl_token_updateCmd)
	addServerFlags(acl_token_updateCmd)

	acl_token_updateCmd.Flags().String("description", "", "A description of the token")
	acl_token_updateCmd.Flags().String("format", "", "Output format.")
	acl_token_updateCmd.Flags().String("id", "", "The Accessor ID of the token to update.")
	acl_token_updateCmd.Flags().Bool("merge-node-identities", false, "Merge the new node identities with the existing node identities")
	acl_token_updateCmd.Flags().Bool("merge-policies", false, "Merge the new policies with the existing policies")
	acl_token_updateCmd.Flags().Bool("merge-roles", false, "Merge the new roles with the existing roles")
	acl_token_updateCmd.Flags().Bool("merge-service-identities", false, "Merge the new service identities with the existing service identities")
	acl_token_updateCmd.Flags().Bool("meta", false, "Indicates that token metadata such as the content hash and raft indices should be shown for each entry")
	acl_token_updateCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_token_updateCmd.Flags().String("node-identity", "", "Name of a node identity to use for this token.")
	acl_token_updateCmd.Flags().String("policy-id", "", "ID of a policy to use for this token.")
	acl_token_updateCmd.Flags().String("policy-name", "", "Name of a policy to use for this token.")
	acl_token_updateCmd.Flags().String("role-id", "", "ID of a role to use for this token.")
	acl_token_updateCmd.Flags().String("role-name", "", "Name of a role to use for this token.")
	acl_token_updateCmd.Flags().String("service-identity", "", "Name of a service identity to use for this token.")
	acl_token_updateCmd.Flags().Bool("upgrade-legacy", false, "Add new polices to a legacy token replacing all existing rules.")
	acl_tokenCmd.AddCommand(acl_token_updateCmd)

	carapace.Gen(acl_token_updateCmd).FlagCompletion(carapace.ActionMap{
		"format":           carapace.ActionValues("pretty", "json"),
		"node-identity":    action.ActionNodeIdentity(acl_token_updateCmd),
		"service-identity": action.ActionServiceIdentity(acl_token_updateCmd),
	})
}
