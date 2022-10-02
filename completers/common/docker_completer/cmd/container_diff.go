package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Inspect changes to files or directories on a container's filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_diffCmd).Standalone()

	containerCmd.AddCommand(container_diffCmd)

	rootAlias(container_diffCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(docker.ActionContainers())
	})
}
