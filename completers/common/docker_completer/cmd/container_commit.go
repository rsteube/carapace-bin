package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a new image from a container's changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_commitCmd).Standalone()

	container_commitCmd.Flags().StringP("author", "a", "", "Author (e.g., \"John Hannibal Smith <hannibal@a-team.com>\")")
	container_commitCmd.Flags().StringP("change", "c", "", "Apply Dockerfile instruction to the created image")
	container_commitCmd.Flags().StringP("message", "m", "", "Commit message")
	container_commitCmd.Flags().BoolP("pause", "p", false, "Pause container during commit (default true)")
	containerCmd.AddCommand(container_commitCmd)

	rootAlias(container_commitCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"change": carapace.ActionFiles(),
		})

		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionContainers(),
			docker.ActionRepositoryTags(),
		)
	})
}
