package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listX11KeymapLayoutsCmd = &cobra.Command{
	Use:   "list-x11-keymap-layouts",
	Short: "Show known X11 keyboard mapping layouts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listX11KeymapLayoutsCmd).Standalone()

	rootCmd.AddCommand(listX11KeymapLayoutsCmd)
}
