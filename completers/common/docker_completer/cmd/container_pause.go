package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause all processes within one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_pauseCmd).Standalone()

	containerCmd.AddCommand(container_pauseCmd)

	rootAlias(container_pauseCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionContainers())
	})
}
