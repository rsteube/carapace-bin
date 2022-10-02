package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var server_bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap the server and retrieve the initial auth token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_bootstrapCmd).Standalone()

	server_bootstrapCmd.Flags().String("context-create", "", "Create a CLI context for this bootstrapped server.")
	server_bootstrapCmd.Flags().Bool("context-set-default", false, "Set the newly bootstrapped server as the default CLI context.")

	addGlobalOptions(server_bootstrapCmd)
	addConnectionOptions(server_bootstrapCmd)

	serverCmd.AddCommand(server_bootstrapCmd)
}
