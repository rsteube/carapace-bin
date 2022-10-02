package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var network_disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect a container from a network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_disconnectCmd).Standalone()

	network_disconnectCmd.Flags().BoolP("force", "f", false, "Force the container to disconnect from a network")
	networkCmd.AddCommand(network_disconnectCmd)

	carapace.Gen(network_disconnectCmd).PositionalCompletion(
		docker.ActionNetworks(),
		docker.ActionContainers(),
	)
}
