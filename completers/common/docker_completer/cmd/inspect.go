package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Return low-level information on Docker objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()

	inspectCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	inspectCmd.Flags().BoolP("size", "s", false, "Display total file sizes if the type is container")
	inspectCmd.Flags().String("type", "", "Return JSON for specified type")
	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).PositionalAnyCompletion(
		carapace.Batch(
			docker.ActionContainers(),
			docker.ActionNetworks(),
			docker.ActionNodes().Suppress("This node is not a swarm manager"),
			docker.ActionRepositoryTags(),
			docker.ActionSecrets().Suppress("This node is not a swarm manager"),
			docker.ActionServices().Suppress("This node is not a swarm manager"),
			docker.ActionVolumes(),
		).ToA(),
	)
}
