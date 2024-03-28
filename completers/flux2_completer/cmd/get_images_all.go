package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_images_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all image statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_images_allCmd).Standalone()
	get_imagesCmd.AddCommand(get_images_allCmd)
}
