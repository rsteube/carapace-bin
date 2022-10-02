package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var acl_token_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an ACL token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_token_createCmd).Standalone()
	addClientFlags(acl_token_createCmd)
	addServerFlags(acl_token_createCmd)

	acl_token_createCmd.Flags().String("accessor", "", "Create the token with this Accessor ID.")
	acl_token_createCmd.Flags().String("description", "", "A description of the token")
	acl_token_createCmd.Flags().String("expires-ttl", "", "Duration of time this token should be valid for")
	acl_token_createCmd.Flags().String("format", "", "Output format.")
	acl_token_createCmd.Flags().Bool("local", false, "Create this as a datacenter local token")
	acl_token_createCmd.Flags().Bool("meta", false, "Indicates that token metadata such as the content hash and raft indices should be shown for each entry")
	acl_token_createCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_token_createCmd.Flags().String("node-identity", "", "Name of a node identity to use for this token.")
	acl_token_createCmd.Flags().String("policy-id", "", "ID of a policy to use for this token.")
	acl_token_createCmd.Flags().String("policy-name", "", "Name of a policy to use for this token.")
	acl_token_createCmd.Flags().String("role-id", "", "ID of a role to use for this token.")
	acl_token_createCmd.Flags().String("role-name", "", "Name of a role to use for this token.")
	acl_token_createCmd.Flags().String("secret", "", "Create the token with this Secret ID.")
	acl_token_createCmd.Flags().String("service-identity", "", "Name of a service identity to use for this token.")
	acl_tokenCmd.AddCommand(acl_token_createCmd)

	carapace.Gen(acl_token_createCmd).FlagCompletion(carapace.ActionMap{
		"format":           carapace.ActionValues("pretty", "json"),
		"node-identity":    action.ActionNodeIdentity(acl_token_createCmd),
		"service-identity": action.ActionServiceIdentity(acl_token_createCmd),
	})
}
