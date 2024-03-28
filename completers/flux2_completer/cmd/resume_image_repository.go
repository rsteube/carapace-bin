package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_image_repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Resume a suspended ImageRepository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_image_repositoryCmd).Standalone()
	resume_imageCmd.AddCommand(resume_image_repositoryCmd)
}
