package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_authMethodCmd = &cobra.Command{
	Use:   "auth-method",
	Short: " Manage Consul's ACL auth methods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_authMethodCmd).Standalone()

	aclCmd.AddCommand(acl_authMethodCmd)
}
