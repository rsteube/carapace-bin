package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var general_permissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: "Show caller permissions for authenticated operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(general_permissionsCmd).Standalone()

	generalCmd.AddCommand(general_permissionsCmd)
}
