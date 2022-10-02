package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snapshot_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a snapshot taken previously with snapshot save",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_deleteCmd).Standalone()

	snapshotCmd.AddCommand(snapshot_deleteCmd)

	carapace.Gen(snapshot_deleteCmd).PositionalCompletion(
		action.ActionLocalMachines(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSnapshots(c.Args[0])
		}),
	)
}
