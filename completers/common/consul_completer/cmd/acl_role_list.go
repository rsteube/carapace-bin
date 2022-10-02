package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_role_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists ACL roles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_role_listCmd).Standalone()
	addClientFlags(acl_role_listCmd)
	addServerFlags(acl_role_listCmd)

	acl_role_listCmd.Flags().String("format", "", "Output format.")
	acl_role_listCmd.Flags().Bool("meta", false, "Indicates that policy metadata such as the content hash and raft indices should be shown for each entry")
	acl_role_listCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_roleCmd.AddCommand(acl_role_listCmd)

	carapace.Gen(acl_role_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
