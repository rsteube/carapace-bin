package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getDesktopViewportCmd = &cobra.Command{
	Use:   "get_desktop_viewport",
	Short: "Report the current viewport's position",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getDesktopViewportCmd).Standalone()
	getDesktopViewportCmd.Flags().Bool("shell", false, "output friendly to shell eval")

	rootCmd.AddCommand(getDesktopViewportCmd)
}
