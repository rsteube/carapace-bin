package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snapshot_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Take a snapshot of the current state of the machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_saveCmd).Standalone()

	snapshot_saveCmd.Flags().BoolP("force", "f", false, "Replace snapshot without confirmation")
	snapshotCmd.AddCommand(snapshot_saveCmd)

	carapace.Gen(snapshot_saveCmd).PositionalCompletion(
		action.ActionLocalMachines(),
	)
}
