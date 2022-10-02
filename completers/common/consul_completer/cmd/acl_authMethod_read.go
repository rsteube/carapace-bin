package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_authMethod_readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read an ACL auth method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_authMethod_readCmd).Standalone()
	addClientFlags(acl_authMethod_readCmd)
	addServerFlags(acl_authMethod_readCmd)

	acl_authMethod_readCmd.Flags().String("format", "", "Output format.")
	acl_authMethod_readCmd.Flags().Bool("meta", false, "Indicates that auth method metadata such as the raft indices should be shown for each entry.")
	acl_authMethod_readCmd.Flags().String("name", "", "The name of the auth method to read.")
	acl_authMethod_readCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_authMethodCmd.AddCommand(acl_authMethod_readCmd)

	carapace.Gen(acl_authMethod_readCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
