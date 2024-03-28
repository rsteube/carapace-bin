package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_helmreleaseCmd = &cobra.Command{
	Use:   "helmrelease",
	Short: "Delete a HelmRelease resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_helmreleaseCmd).Standalone()
	deleteCmd.AddCommand(delete_helmreleaseCmd)
}
