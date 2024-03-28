package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_images_repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Get ImageRepository status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_images_repositoryCmd).Standalone()
	get_imagesCmd.AddCommand(get_images_repositoryCmd)
}
