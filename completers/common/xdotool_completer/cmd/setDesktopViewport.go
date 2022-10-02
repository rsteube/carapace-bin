package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setDesktopViewportCmd = &cobra.Command{
	Use:   "set_desktop_viewport",
	Short: "Move the viewport to the given position",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setDesktopViewportCmd).Standalone()

	rootCmd.AddCommand(setDesktopViewportCmd)
}
