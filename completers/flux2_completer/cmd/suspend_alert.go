package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Suspend reconciliation of Alert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_alertCmd).Standalone()
	suspendCmd.AddCommand(suspend_alertCmd)
}
