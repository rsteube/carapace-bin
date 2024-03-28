package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_source_helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Reconcile a HelmRepository source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_source_helmCmd).Standalone()
	reconcile_sourceCmd.AddCommand(reconcile_source_helmCmd)
}
