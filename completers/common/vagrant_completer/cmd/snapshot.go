package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "manages snapshots: saving, restoring, etc.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshotCmd).Standalone()

	rootCmd.AddCommand(snapshotCmd)
}
