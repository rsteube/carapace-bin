package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_bindingRule_readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read an ACL binding rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_bindingRule_readCmd).Standalone()
	addClientFlags(acl_bindingRule_readCmd)
	addServerFlags(acl_bindingRule_readCmd)

	acl_bindingRule_readCmd.Flags().String("format", "", "Output format.")
	acl_bindingRule_readCmd.Flags().String("id", "", "The ID of the binding rule to read.")
	acl_bindingRule_readCmd.Flags().Bool("meta", false, "Indicates that binding rule metadata such as the raft indices should be shown for each entry.")
	acl_bindingRule_readCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_bindingRuleCmd.AddCommand(acl_bindingRule_readCmd)

	carapace.Gen(acl_bindingRule_readCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
