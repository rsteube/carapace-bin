package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_receiverCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Suspend reconciliation of Receiver",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_receiverCmd).Standalone()
	suspendCmd.AddCommand(suspend_receiverCmd)
}
