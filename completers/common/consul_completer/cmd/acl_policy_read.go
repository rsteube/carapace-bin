package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_policy_readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read an ACL policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_policy_readCmd).Standalone()
	addClientFlags(acl_policy_readCmd)
	addServerFlags(acl_policy_readCmd)

	acl_policy_readCmd.Flags().String("format", "", "Output format.")
	acl_policy_readCmd.Flags().String("id", "", "The ID of the policy to read.")
	acl_policy_readCmd.Flags().Bool("meta", false, "Indicates that policy metadata such as the content hash and raft indices should be shown for each entry")
	acl_policy_readCmd.Flags().String("name", "", "The name of the policy to read.")
	acl_policy_readCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_policyCmd.AddCommand(acl_policy_readCmd)

	carapace.Gen(acl_policy_readCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
