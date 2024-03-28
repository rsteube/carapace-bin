package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_image_repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Suspend reconciliation of an ImageRepository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_image_repositoryCmd).Standalone()
	suspend_imageCmd.AddCommand(suspend_image_repositoryCmd)
}
