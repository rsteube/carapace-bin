package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Saves, restores and inspects snapshots of Consul server state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshotCmd).Standalone()

	rootCmd.AddCommand(snapshotCmd)
}
