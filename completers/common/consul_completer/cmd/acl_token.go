package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage Consul's ACL tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_tokenCmd).Standalone()

	aclCmd.AddCommand(acl_tokenCmd)
}
