package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_image_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Delete an ImageUpdateAutomation object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_image_updateCmd).Standalone()
	delete_imageCmd.AddCommand(delete_image_updateCmd)
}
