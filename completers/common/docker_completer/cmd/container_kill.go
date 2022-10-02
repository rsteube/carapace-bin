package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill one or more running containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_killCmd).Standalone()

	container_killCmd.Flags().StringP("signal", "s", "", "Signal to send to the container (default \"KILL\")")
	containerCmd.AddCommand(container_killCmd)

	rootAlias(container_killCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"signal": ps.ActionKillSignals(),
		})

		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionContainers())
	})
}
