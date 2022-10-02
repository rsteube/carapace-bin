package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var acl_role_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an ACL role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_role_updateCmd).Standalone()
	addClientFlags(acl_role_updateCmd)
	addServerFlags(acl_role_updateCmd)

	acl_role_updateCmd.Flags().String("description", "", "A description of the role")
	acl_role_updateCmd.Flags().String("format", "", "Output format.")
	acl_role_updateCmd.Flags().String("id", "", "The ID of the role to update.")
	acl_role_updateCmd.Flags().Bool("meta", false, "Indicates that role metadata such as the content hash and raft indices should be shown for each entry")
	acl_role_updateCmd.Flags().String("name", "", "The role name.")
	acl_role_updateCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_role_updateCmd.Flags().Bool("no-merge", false, "Do not merge the current role information with what is provided to the command.")
	acl_role_updateCmd.Flags().String("node-identity", "", "Name of a node identity to use for this role.")
	acl_role_updateCmd.Flags().String("policy-id", "", "ID of a policy to use for this role.")
	acl_role_updateCmd.Flags().String("policy-name", "", "Name of a policy to use for this role.")
	acl_role_updateCmd.Flags().String("service-identity", "", "Name of a service identity to use for this role.")
	acl_roleCmd.AddCommand(acl_role_updateCmd)

	carapace.Gen(acl_role_createCmd).FlagCompletion(carapace.ActionMap{
		"format":           carapace.ActionValues("pretty", "json"),
		"node-identity":    action.ActionNodeIdentity(acl_role_createCmd),
		"service-identity": action.ActionServiceIdentity(acl_role_createCmd),
	})
}
