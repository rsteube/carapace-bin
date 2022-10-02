package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listX11KeymapModelsCmd = &cobra.Command{
	Use:   "list-x11-keymap-models",
	Short: "Show known X11 keyboard mapping models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listX11KeymapModelsCmd).Standalone()

	rootCmd.AddCommand(listX11KeymapModelsCmd)
}
