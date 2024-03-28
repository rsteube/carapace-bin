package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_image_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Resume a suspended ImageUpdateAutomation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_image_updateCmd).Standalone()
	resume_imageCmd.AddCommand(resume_image_updateCmd)
}
