package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var aclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Interact with Consul's ACLs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aclCmd).Standalone()

	rootCmd.AddCommand(aclCmd)
}
