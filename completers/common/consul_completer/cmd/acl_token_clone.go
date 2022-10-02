package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone an ACL token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_cloneCmd).Standalone()
	addClientFlags(acl_cloneCmd)
	addServerFlags(acl_cloneCmd)

	acl_cloneCmd.Flags().String("description", "", "A description of the new cloned token")
	acl_cloneCmd.Flags().String("format", "", "Output format.")
	acl_cloneCmd.Flags().String("id", "", "The Accessor ID of the token to clone.")
	acl_cloneCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	acl_tokenCmd.AddCommand(acl_cloneCmd)

	carapace.Gen(acl_cloneCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
