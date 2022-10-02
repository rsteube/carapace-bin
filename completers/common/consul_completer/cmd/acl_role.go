package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_roleCmd = &cobra.Command{
	Use:   "role",
	Short: "Manage Consul's ACL roles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_roleCmd).Standalone()

	aclCmd.AddCommand(acl_roleCmd)
}
