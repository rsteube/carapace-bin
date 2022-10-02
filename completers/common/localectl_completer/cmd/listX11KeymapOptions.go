package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listX11KeymapOptionsCmd = &cobra.Command{
	Use:   "list-x11-keymap-options",
	Short: "Show known X11 keyboard mapping options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listX11KeymapOptionsCmd).Standalone()

	rootCmd.AddCommand(listX11KeymapOptionsCmd)
}
