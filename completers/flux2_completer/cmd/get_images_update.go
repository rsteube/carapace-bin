package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_images_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Get ImageUpdateAutomation status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_images_updateCmd).Standalone()
	get_imagesCmd.AddCommand(get_images_updateCmd)
}
