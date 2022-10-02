package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_role_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an ACL role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_role_deleteCmd).Standalone()
	addClientFlags(acl_role_deleteCmd)
	addServerFlags(acl_role_deleteCmd)

	acl_role_deleteCmd.Flags().String("id", "", "The ID of the role to delete.")
	acl_role_deleteCmd.Flags().String("name", "", "The name of the role to delete.")
	acl_role_deleteCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_roleCmd.AddCommand(acl_role_deleteCmd)
}
