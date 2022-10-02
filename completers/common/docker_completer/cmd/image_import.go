package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import the contents from a tarball to create a filesystem image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_importCmd).Standalone()

	image_importCmd.Flags().StringP("change", "c", "", "Apply Dockerfile instruction to the created image")
	image_importCmd.Flags().StringP("message", "m", "", "Set commit message for imported image")
	image_importCmd.Flags().String("platform", "", "Set platform if server is multi-platform capable")
	imageCmd.AddCommand(image_importCmd)

	rootAlias(image_importCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			carapace.ActionFiles(),
			docker.ActionRepositoryTags(),
		)
	})
}
