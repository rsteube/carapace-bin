package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_alertProviderCmd = &cobra.Command{
	Use:   "alert-provider",
	Short: "Delete a Provider resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_alertProviderCmd).Standalone()
	deleteCmd.AddCommand(delete_alertProviderCmd)
}
