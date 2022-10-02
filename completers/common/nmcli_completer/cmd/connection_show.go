package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_showCmd = &cobra.Command{
	Use:   "show",
	Short: "show details for specified connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_showCmd).Standalone()

	connection_showCmd.Flags().Bool("active", false, "show active connections")
	connection_showCmd.Flags().String("order", "", "specify order")

	connectionCmd.AddCommand(connection_showCmd)

	carapace.Gen(connection_showCmd).PositionalAnyCompletion(net.ActionConnections())
}
