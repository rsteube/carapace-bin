package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_bindingRule_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an ACL binding rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_bindingRule_updateCmd).Standalone()
	addClientFlags(acl_bindingRule_updateCmd)
	addServerFlags(acl_bindingRule_updateCmd)

	acl_bindingRule_updateCmd.Flags().String("bind-name", "", "Name to bind on match.")
	acl_bindingRule_updateCmd.Flags().String("bind-type", "", "Type of binding to perform.")
	acl_bindingRule_updateCmd.Flags().String("description", "", "A description of the binding rule.")
	acl_bindingRule_updateCmd.Flags().String("format", "", "Output format.")
	acl_bindingRule_updateCmd.Flags().String("id", "", "The ID of the binding rule to update.")
	acl_bindingRule_updateCmd.Flags().Bool("meta", false, "Indicates that binding rule metadata such as the raft indices should be shown for each entry.")
	acl_bindingRule_updateCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_bindingRule_updateCmd.Flags().Bool("no-merge", false, "Do not merge the current binding rule information with what is provided to the command.")
	acl_bindingRule_updateCmd.Flags().String("selector", "", "Selector is an expression that matches against verified identity attributes returned from the auth method during login.")
	acl_bindingRuleCmd.AddCommand(acl_bindingRule_updateCmd)

	carapace.Gen(acl_bindingRule_updateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
