package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_token_readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read an ACL token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_token_readCmd).Standalone()
	addClientFlags(acl_token_readCmd)
	addServerFlags(acl_token_readCmd)

	acl_token_readCmd.Flags().String("format", "", "Output format.")
	acl_token_readCmd.Flags().String("id", "", "The Accessor ID of the token to read.")
	acl_token_readCmd.Flags().Bool("meta", false, "Indicates that token metadata such as the content hash and Raft indices should be shown for each entry")
	acl_token_readCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_token_readCmd.Flags().Bool("self", false, "Indicates that the current HTTP token should be read by secret ID")
	acl_tokenCmd.AddCommand(acl_token_readCmd)

	carapace.Gen(acl_token_readCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
