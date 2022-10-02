package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snapshot_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snapshots taken for a machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_listCmd).Standalone()

	snapshotCmd.AddCommand(snapshot_listCmd)

	carapace.Gen(snapshot_listCmd).PositionalCompletion(
		action.ActionLocalMachines(),
	)
}
