package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Reconcile an Alert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_alertCmd).Standalone()
	reconcileCmd.AddCommand(reconcile_alertCmd)
}
