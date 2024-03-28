package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_image_repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Delete an ImageRepository object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_image_repositoryCmd).Standalone()
	delete_imageCmd.AddCommand(delete_image_repositoryCmd)
}
