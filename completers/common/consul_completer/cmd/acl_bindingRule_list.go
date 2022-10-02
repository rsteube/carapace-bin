package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_bindingRule_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists ACL binding rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_bindingRule_listCmd).Standalone()
	addClientFlags(acl_bindingRule_listCmd)
	addServerFlags(acl_bindingRule_listCmd)

	acl_bindingRule_listCmd.Flags().String("format", "", "Output format.")
	acl_bindingRule_listCmd.Flags().Bool("meta", false, "Indicates that binding rule metadata such as the raft indices should be shown for each entry.")
	acl_bindingRule_listCmd.Flags().String("method", "", "Only show rules linked to the auth method with the given name.")
	acl_bindingRule_listCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_bindingRuleCmd.AddCommand(acl_bindingRule_listCmd)

	carapace.Gen(acl_bindingRule_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
