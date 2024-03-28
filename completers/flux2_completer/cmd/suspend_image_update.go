package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_image_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Suspend reconciliation of an ImageUpdateAutomation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_image_updateCmd).Standalone()
	suspend_imageCmd.AddCommand(suspend_image_updateCmd)
}
