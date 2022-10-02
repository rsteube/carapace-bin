package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var org_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all users in an org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(org_lsCmd).Standalone()

	orgCmd.AddCommand(org_lsCmd)
}
