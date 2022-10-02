package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_role_readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read an ACL role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_role_readCmd).Standalone()
	addClientFlags(acl_role_readCmd)
	addServerFlags(acl_role_readCmd)

	acl_role_readCmd.Flags().String("format", "", "Output format.")
	acl_role_readCmd.Flags().String("id", "", "The ID of the role to read.")
	acl_role_readCmd.Flags().Bool("meta", false, "Indicates that role metadata such as the content hash and raft indices should be shown for each entry")
	acl_role_readCmd.Flags().String("name", "", "The name of the role to read.")
	acl_role_readCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_roleCmd.AddCommand(acl_role_readCmd)

	carapace.Gen(acl_role_readCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
