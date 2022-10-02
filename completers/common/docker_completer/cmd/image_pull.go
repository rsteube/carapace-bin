package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull an image or a repository from a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pullCmd).Standalone()

	image_pullCmd.Flags().BoolP("all-tags", "a", false, "Download all tagged images in the repository")
	image_pullCmd.Flags().Bool("disable-content-trust", false, "Skip image verification (default true)")
	image_pullCmd.Flags().String("platform", "", "Set platform if server is multi-platform capable")
	image_pullCmd.Flags().BoolP("quiet", "q", false, "Suppress verbose output")
	imageCmd.AddCommand(image_pullCmd)

	rootAlias(image_pullCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(docker.ActionRepositoryTags())
	})
}
