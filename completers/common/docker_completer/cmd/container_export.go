package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a container's filesystem as a tar archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_exportCmd).Standalone()

	container_exportCmd.Flags().StringP("output", "o", "", "Write to a file, instead of STDOUT")
	containerCmd.AddCommand(container_exportCmd)

	rootAlias(container_exportCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionContainers(),
		)
	})
}
