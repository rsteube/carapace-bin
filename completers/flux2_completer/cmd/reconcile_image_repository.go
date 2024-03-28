package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_image_repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Reconcile an ImageRepository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_image_repositoryCmd).Standalone()
	reconcile_imageCmd.AddCommand(reconcile_image_repositoryCmd)
}
