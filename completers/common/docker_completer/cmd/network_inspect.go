package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var network_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display detailed information on one or more networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_inspectCmd).Standalone()

	network_inspectCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	network_inspectCmd.Flags().BoolP("verbose", "v", false, "Verbose output for diagnostics")
	networkCmd.AddCommand(network_inspectCmd)

	carapace.Gen(network_inspectCmd).PositionalAnyCompletion(docker.ActionNetworks())
}
