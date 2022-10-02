package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var acl_policy_updateCmd = &cobra.Command{
	Use:   "update",
	Short: " Update an ACL policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_policy_updateCmd).Standalone()
	addClientFlags(acl_policy_updateCmd)
	addServerFlags(acl_policy_updateCmd)

	acl_policy_updateCmd.Flags().String("description", "", "A description of the policy")
	acl_policy_updateCmd.Flags().String("format", "", "Output format.")
	acl_policy_updateCmd.Flags().String("id", "", "The ID of the policy to update.")
	acl_policy_updateCmd.Flags().Bool("meta", false, "Indicates that policy metadata such as the content hash and raft indices should be shown for each entry")
	acl_policy_updateCmd.Flags().String("name", "", "The policies name.")
	acl_policy_updateCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_policy_updateCmd.Flags().Bool("no-merge", false, "Do not merge the current policy information with what is provided to the command.")
	acl_policy_updateCmd.Flags().String("rules", "", "The policy rules.")
	acl_policy_updateCmd.Flags().String("valid-datacenter", "", "Datacenter that the policy should be valid within.")
	acl_policyCmd.AddCommand(acl_policy_updateCmd)

	carapace.Gen(acl_policy_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
		"rules":  action.ActionOptionalFiles(),
	})
}
