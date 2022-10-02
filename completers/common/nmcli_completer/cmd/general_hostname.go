package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var general_hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Get or change persistent system hostname",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(general_hostnameCmd).Standalone()

	generalCmd.AddCommand(general_hostnameCmd)

	carapace.Gen(general_hostnameCmd).PositionalCompletion(
		net.ActionHosts(),
	)
}
