package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_alertProviderCmd = &cobra.Command{
	Use:   "alert-provider",
	Short: "Reconcile a Provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_alertProviderCmd).Standalone()
	reconcileCmd.AddCommand(reconcile_alertProviderCmd)
}
