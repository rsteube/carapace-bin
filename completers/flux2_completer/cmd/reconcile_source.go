package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Reconcile sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_sourceCmd).Standalone()
	reconcileCmd.AddCommand(reconcile_sourceCmd)
}
