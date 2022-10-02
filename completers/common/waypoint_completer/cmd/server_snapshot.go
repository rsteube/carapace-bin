package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var server_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Write a backup of the server data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_snapshotCmd).Standalone()

	addGlobalOptions(server_snapshotCmd)

	serverCmd.AddCommand(server_snapshotCmd)
}
