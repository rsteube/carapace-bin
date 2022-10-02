package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Display a live stream of container(s) resource usage statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_statsCmd).Standalone()

	container_statsCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	container_statsCmd.Flags().String("format", "", "Pretty-print images using a Go template")
	container_statsCmd.Flags().Bool("no-stream", false, "Disable streaming stats and only pull the first result")
	container_statsCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	containerCmd.AddCommand(container_statsCmd)

	rootAlias(container_statsCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionContainers())
	})
}
