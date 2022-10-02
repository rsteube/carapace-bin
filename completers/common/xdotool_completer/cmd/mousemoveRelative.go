package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mousemoveRelativeCmd = &cobra.Command{
	Use:   "mousemove_relative",
	Short: "Move the mouse x,y pixels relative to the current position of the mouse cursor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mousemoveRelativeCmd).Standalone()

	mousemoveRelativeCmd.Flags().Bool("clearmodifiers", false, "See CLEARMODIFIERS")
	mousemoveRelativeCmd.Flags().Bool("polar", false, "Use polar coordinates")
	mousemoveRelativeCmd.Flags().Bool("sync", false, "After sending the mouse move request, wait until the mouse is actually moved")
	rootCmd.AddCommand(mousemoveRelativeCmd)
}
