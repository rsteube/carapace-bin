package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Block until one or more containers stop, then print their exit codes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_waitCmd).Standalone()

	containerCmd.AddCommand(container_waitCmd)

	rootAlias(container_waitCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionContainers())
	})
}
