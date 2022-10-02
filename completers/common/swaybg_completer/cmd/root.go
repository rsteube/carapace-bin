package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swaybg",
	Short: "Background for Wayland",
	Long:  "https://github.com/swaywm/swaybg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("color", "c", "", "Set the background color.")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message and quit.")
	rootCmd.Flags().StringP("image", "i", "", "Set the image to display.")
	rootCmd.Flags().StringP("mode", "m", "", "Set the mode to use for the image.")
	rootCmd.Flags().StringP("output", "o", "", "Set the output to operate on or * for all.")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version number and quit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":  color.ActionHexColors(),
		"image":  carapace.ActionFiles(),
		"mode":   carapace.ActionValues("stretch", "fill", "fit", "center", "tile", "solid_color"),
		"output": carapace.ActionFiles(),
	})
}
