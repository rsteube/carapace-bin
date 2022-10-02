package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_authMethod_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an ACL auth method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_authMethod_deleteCmd).Standalone()

	acl_authMethod_deleteCmd.Flags().String("name", "", "The name of the auth method to delete.")
	acl_authMethod_deleteCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_authMethodCmd.AddCommand(acl_authMethod_deleteCmd)
}
