package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_token_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List ACL tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_token_listCmd).Standalone()
	addClientFlags(acl_token_listCmd)
	addServerFlags(acl_token_listCmd)

	acl_token_listCmd.Flags().String("format", "", "Output format.")
	acl_token_listCmd.Flags().Bool("meta", false, "Indicates that token metadata such as the content hash and Raft indices should be shown for each entry")
	acl_token_listCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_tokenCmd.AddCommand(acl_token_listCmd)

	carapace.Gen(acl_token_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
