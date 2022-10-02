package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snapshot_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Take a snapshot of the current state of the machine and 'push' it onto the stack of states",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_pushCmd).Standalone()

	snapshotCmd.AddCommand(snapshot_pushCmd)

	carapace.Gen(snapshot_pushCmd).PositionalCompletion(
		action.ActionLocalMachines(),
	)
}
