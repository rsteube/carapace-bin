package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var windowunmapCmd = &cobra.Command{
	Use:   "windowunmap",
	Short: "Unmap a window, making it no longer appear on your screen",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowunmapCmd).Standalone()

	windowunmapCmd.Flags().Bool("sync", false, "After requesting the window unmap, wait until the window is actually unmapped (hidden)")
	rootCmd.AddCommand(windowunmapCmd)
}
