package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_topCmd = &cobra.Command{
	Use:   "top",
	Short: "Display the running processes of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_topCmd).Standalone()

	containerCmd.AddCommand(container_topCmd)

	rootAlias(container_topCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionContainers())
	})
}
