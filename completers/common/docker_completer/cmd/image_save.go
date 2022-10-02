package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save one or more images to a tar archive (streamed to STDOUT by default)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_saveCmd).Standalone()

	image_saveCmd.Flags().StringP("output", "o", "", "Write to a file, instead of STDOUT")
	imageCmd.AddCommand(image_saveCmd)

	rootAlias(image_saveCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"output": carapace.ActionFiles(),
		})

		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionRepositoryTags())
	})
}
