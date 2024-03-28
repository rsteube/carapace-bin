package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_helmreleaseCmd = &cobra.Command{
	Use:   "helmrelease",
	Short: "Reconcile a HelmRelease resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_helmreleaseCmd).Standalone()
	reconcile_helmreleaseCmd.Flags().Bool("with-source", false, "reconcile HelmRelease source")
	reconcileCmd.AddCommand(reconcile_helmreleaseCmd)
}
