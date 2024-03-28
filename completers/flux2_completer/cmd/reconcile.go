package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcileCmd = &cobra.Command{
	Use:   "reconcile",
	Short: "Reconcile sources and resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcileCmd).Standalone()
	rootCmd.AddCommand(reconcileCmd)
}
