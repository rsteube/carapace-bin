package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var snapshot_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Saves snapshot of Consul server stat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_saveCmd).Standalone()
	addClientFlags(snapshot_saveCmd)
	addServerFlags(snapshot_saveCmd)

	snapshotCmd.AddCommand(snapshot_saveCmd)

	carapace.Gen(snapshot_saveCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
