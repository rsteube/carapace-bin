package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var org_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a user in an org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(org_rmCmd).Standalone()

	orgCmd.AddCommand(org_rmCmd)
}
