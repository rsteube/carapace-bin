package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_importCmd = &cobra.Command{
	Use:     "import [IMAGE | ARCHIVE [IMAGE | ARCHIVE...]]",
	Short:   "Import image(s) from docker into k3d cluster(s).",
	Aliases: []string{"load"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_importCmd).Standalone()

	image_importCmd.Flags().StringSliceP("cluster", "c", []string{}, "Select clusters to load the image to.")
	image_importCmd.Flags().BoolP("keep-tarball", "k", false, "Do not delete the tarball containing the saved images from the shared volume")
	image_importCmd.Flags().BoolP("keep-tools", "t", false, "Do not delete the tools node after import")
	image_importCmd.Flags().StringP("mode", "m", "", "Which method to use to import images into the cluster [auto, direct, tools]. See https://k3d.io/stable/usage/importing_images/")
	imageCmd.AddCommand(image_importCmd)
}
