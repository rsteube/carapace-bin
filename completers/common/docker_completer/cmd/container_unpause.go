package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "Unpause all processes within one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_unpauseCmd).Standalone()

	containerCmd.AddCommand(container_unpauseCmd)

	rootAlias(container_unpauseCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionContainers())
	})
}
