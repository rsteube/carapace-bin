package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_receiverCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Reconcile a Receiver",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_receiverCmd).Standalone()
	reconcileCmd.AddCommand(reconcile_receiverCmd)
}
