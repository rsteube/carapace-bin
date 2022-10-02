package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var server_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Stage a snapshot on the current server for data restoration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_restoreCmd).Standalone()

	server_restoreCmd.Flags().Bool("exit", false, "After restoring, the server should exit so it can be restarted.")
	serverCmd.AddCommand(server_restoreCmd)
}
