package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_image_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Reconcile an ImageUpdateAutomation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_image_updateCmd).Standalone()
	reconcile_imageCmd.AddCommand(reconcile_image_updateCmd)
}
