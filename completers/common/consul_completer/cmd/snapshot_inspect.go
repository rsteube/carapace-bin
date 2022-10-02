package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var snapshot_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Displays information about a Consul snapshot file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_inspectCmd).Standalone()

	snapshotCmd.AddCommand(snapshot_inspectCmd)

	carapace.Gen(snapshot_inspectCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
