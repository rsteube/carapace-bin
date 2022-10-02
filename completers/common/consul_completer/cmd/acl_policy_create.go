package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var acl_policy_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an ACL policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_policy_createCmd).Standalone()
	addClientFlags(acl_policy_createCmd)
	addServerFlags(acl_policy_createCmd)

	acl_policy_createCmd.Flags().String("format", "", "Output format.")
	acl_policy_createCmd.Flags().String("from-token", "", "The legacy token to retrieve the rules for when creating this policy.")
	acl_policy_createCmd.Flags().Bool("meta", false, "Indicates that policy metadata such as the content hash and raft indices should be shown for each entry")
	acl_policy_createCmd.Flags().String("name", "", "The new policy's name.")
	acl_policy_createCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_policy_createCmd.Flags().String("rules", "", "The policy rules.")
	acl_policy_createCmd.Flags().Bool("token-secret", false, "Indicates the token provided with -from-token is a SecretID and not an AccessorID")
	acl_policy_createCmd.Flags().StringArray("valid-datacenter", []string{}, "Datacenter that the policy should be valid within.")
	acl_policyCmd.AddCommand(acl_policy_createCmd)

	carapace.Gen(acl_policy_createCmd).FlagCompletion(carapace.ActionMap{
		"format":           carapace.ActionValues("pretty", "json"),
		"rules":            action.ActionOptionalFiles(),
		"valid-datacenter": action.ActionDatacenters(acl_policy_createCmd),
	})
}
