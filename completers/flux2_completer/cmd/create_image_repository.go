package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_image_repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Create or update an ImageRepository object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_image_repositoryCmd).Standalone()
	create_image_repositoryCmd.Flags().String("cert-ref", "", "the name of a secret to use for TLS certificates")
	create_image_repositoryCmd.Flags().String("image", "", "the image repository to scan; e.g., library/alpine")
	create_image_repositoryCmd.Flags().String("scan-timeout", "", "a timeout for scanning; this defaults to the interval if not set")
	create_image_repositoryCmd.Flags().String("secret-ref", "", "the name of a docker-registry secret to use for credentials")
	create_imageCmd.AddCommand(create_image_repositoryCmd)
}
