package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_rmCmd).Standalone()

	image_rmCmd.Flags().BoolP("force", "f", false, "Force removal of the image")
	image_rmCmd.Flags().Bool("no-prune", false, "Do not delete untagged parents")
	imageCmd.AddCommand(image_rmCmd)

	rootAlias(image_rmCmd, func(cmd *cobra.Command, isAlias bool) {
		if isAlias {
			cmd.Use = "rmi"
		}
		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionRepositoryTags())
	})
}
