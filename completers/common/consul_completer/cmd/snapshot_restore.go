package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var snapshot_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restores snapshot of Consul server state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_restoreCmd).Standalone()
	addClientFlags(snapshot_restoreCmd)
	addServerFlags(snapshot_restoreCmd)

	snapshotCmd.AddCommand(snapshot_restoreCmd)

	carapace.Gen(snapshot_restoreCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
