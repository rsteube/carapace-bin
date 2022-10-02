package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_authMethod_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists ACL auth methods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_authMethod_listCmd).Standalone()
	addClientFlags(acl_authMethod_listCmd)
	addServerFlags(acl_authMethod_listCmd)

	acl_authMethod_listCmd.Flags().String("format", "", "Output format")
	acl_authMethod_listCmd.Flags().Bool("meta", false, "Indicates that auth method metadata such as the raft indices should be shown for each entry.")
	acl_authMethod_listCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_authMethodCmd.AddCommand(acl_authMethod_listCmd)

	carapace.Gen(acl_authMethod_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
