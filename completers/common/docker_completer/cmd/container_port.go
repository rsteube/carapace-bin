package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_portCmd = &cobra.Command{
	Use:   "port",
	Short: "List port mappings or a specific mapping for the container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_portCmd).Standalone()

	containerCmd.AddCommand(container_portCmd)

	rootAlias(container_portCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionContainers(),
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				return docker.ActionContainerPorts(c.Args[0])
			}),
		)
	})
}
