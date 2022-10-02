package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push an image or a repository to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pushCmd).Standalone()

	image_pushCmd.Flags().Bool("disable-content-trust", false, "Skip image signing (default true)")
	imageCmd.AddCommand(image_pushCmd)

	rootAlias(image_pushCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionRepositoryTags(),
		)
	})
}
