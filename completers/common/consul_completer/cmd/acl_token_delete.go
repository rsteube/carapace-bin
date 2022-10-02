package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_token_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an ACL token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_token_deleteCmd).Standalone()
	addClientFlags(acl_token_deleteCmd)
	addServerFlags(acl_token_deleteCmd)

	acl_token_deleteCmd.Flags().String("id", "", "The Accessor ID of the token to delete.")
	acl_token_deleteCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_tokenCmd.AddCommand(acl_token_deleteCmd)
}
