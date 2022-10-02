package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_exportCmd).Standalone()

	connectionCmd.AddCommand(connection_exportCmd)

	carapace.Gen(connection_exportCmd).PositionalCompletion(
		net.ActionConnections(),
		carapace.ActionFiles(),
	)
}
