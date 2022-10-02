package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap Consul's ACL system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_bootstrapCmd).Standalone()
	addClientFlags(acl_bootstrapCmd)
	addServerFlags(acl_bootstrapCmd)
	acl_bootstrapCmd.Flags().String("format", "", "Output format.")

	aclCmd.AddCommand(acl_bootstrapCmd)

	carapace.Gen(acl_bootstrapCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
