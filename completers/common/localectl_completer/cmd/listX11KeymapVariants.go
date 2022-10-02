package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listX11KeymapVariantsCmd = &cobra.Command{
	Use:   "list-x11-keymap-variants",
	Short: "Show known X11 keyboard mapping variants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listX11KeymapVariantsCmd).Standalone()

	rootCmd.AddCommand(listX11KeymapVariantsCmd)
}
