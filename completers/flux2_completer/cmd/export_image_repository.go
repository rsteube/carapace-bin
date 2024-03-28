package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_image_repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Export ImageRepository resources in YAML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_image_repositoryCmd).Standalone()
	export_imageCmd.AddCommand(export_image_repositoryCmd)
}
