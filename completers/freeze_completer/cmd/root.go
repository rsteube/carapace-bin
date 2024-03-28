package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/chroma"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "freeze",
	Short: "Generate images of code and terminal output",
	Long:  "https://github.com/charmbracelet/freeze",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("background", "b", "", "Apply a background fill")
	rootCmd.Flags().String("border.color", "", "Border color")
	rootCmd.Flags().StringP("border.radius", "r", "", "Corner radius of window")
	rootCmd.Flags().String("border.width", "", "Border width thickness")
	rootCmd.Flags().StringP("config", "c", "", "Base configuration file or template")
	rootCmd.Flags().StringP("execute", "x", "", "Capture output of command execution")
	rootCmd.Flags().String("font.family", "", "Font family to use for code")
	rootCmd.Flags().String("font.file", "", "Font file to embed")
	rootCmd.Flags().Bool("font.ligatures", false, "Use ligatures in the font")
	rootCmd.Flags().String("font.size", "", "Font size to use for code")
	rootCmd.Flags().StringP("height", "H", "", "Height of terminal window")
	rootCmd.Flags().BoolP("interactive", "i", false, "Use an interactive form for configuration options")
	rootCmd.Flags().StringP("language", "l", "", "Language of code file")
	rootCmd.Flags().String("line-height", "", "Line height relative to font size")
	rootCmd.Flags().String("lines", "", "Lines to capture (start,end)")
	rootCmd.Flags().StringP("margin", "m", "", "Apply margin to the window")
	rootCmd.Flags().StringP("output", "o", "", "Output location for .svg, .png, or .webp")
	rootCmd.Flags().StringP("padding", "p", "", "Apply padding to the code")
	rootCmd.Flags().String("shadow.blur", "", "Shadow Gaussian Blur")
	rootCmd.Flags().String("shadow.x", "", "Shadow offset x coordinate")
	rootCmd.Flags().String("shadow.y", "", "Shadow offset y coordinate")
	rootCmd.Flags().Bool("show-line-numbers", false, "Show line numbers")
	rootCmd.Flags().StringP("theme", "t", "", "Theme to use for syntax highlighting")
	rootCmd.Flags().StringP("width", "W", "", "Width of terminal window")
	rootCmd.Flags().Bool("window", false, "Display window controls")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"background": color.ActionHexColors(),
		"config": carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionValuesDescribed(
				"base", "Simple screenshot of code",
				"full", "macOS-like screenshot",
				"user", "Uses ~/.config/freeze/user.json",
			).Tag("default configurations"),
		).ToA(),
		"execute":     bridge.ActionCarapaceBin().Split(),
		"font.family": os.ActionFontFamilies(),
		"font.file":   carapace.ActionFiles(),
		"language":    chroma.ActionLexers(),
		"output":      carapace.ActionFiles(),
		"theme":       chroma.ActionStyles(),
	})
}
