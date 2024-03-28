package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_kustomizationCmd = &cobra.Command{
	Use:   "kustomization",
	Short: "Reconcile a Kustomization resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_kustomizationCmd).Standalone()
	reconcile_kustomizationCmd.Flags().Bool("with-source", false, "reconcile Kustomization source")
	reconcileCmd.AddCommand(reconcile_kustomizationCmd)
}
