package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_bindingRule_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an ACL binding rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_bindingRule_createCmd).Standalone()
	addClientFlags(acl_bindingRule_createCmd)
	addServerFlags(acl_bindingRule_createCmd)

	acl_bindingRule_createCmd.Flags().String("bind-name", "", "Name to bind on match.")
	acl_bindingRule_createCmd.Flags().String("bind-type", "", "Type of binding to perform.")
	acl_bindingRule_createCmd.Flags().String("description", "", "A description of the binding rule.")
	acl_bindingRule_createCmd.Flags().String("format", "", "Output format.")
	acl_bindingRule_createCmd.Flags().Bool("meta", false, "Indicates that binding rule metadata such as the raft indices should be shown for each entry.")
	acl_bindingRule_createCmd.Flags().String("method", "", "The auth method's name for which this binding rule applies.")
	acl_bindingRule_createCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_bindingRule_createCmd.Flags().String("selector", "", "Selector is an expression that matches against verified identity attributes returned from the auth method during login.")
	acl_bindingRuleCmd.AddCommand(acl_bindingRule_createCmd)

	carapace.Gen(acl_bindingRule_createCmd).FlagCompletion(carapace.ActionMap{
		"bind-type": carapace.ActionValues("service", "role"),
		"format":    carapace.ActionValues("pretty", "json"),
	})
}
