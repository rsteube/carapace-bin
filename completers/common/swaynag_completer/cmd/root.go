package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/color"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swaynag",
	Short: "Show a warning or error message with buttons",
	Long:  "https://github.com/swaywm/sway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("background", "", "Background color.")
	rootCmd.Flags().String("border", "", "Border color.")
	rootCmd.Flags().String("border-bottom", "", "Bottom border color.")
	rootCmd.Flags().String("border-bottom-size", "", "Thickness of the bar border.")
	rootCmd.Flags().StringP("button", "b", "", "Create a button with text that executes action in a terminal when pressed.")
	rootCmd.Flags().String("button-background", "", "Button background color.")
	rootCmd.Flags().String("button-border-size", "", "Thickness for the button border.")
	rootCmd.Flags().StringP("button-dismiss", "z", "", "Create a button with text that dismisses swaynag, and executes action in a terminal when pressed.")
	rootCmd.Flags().String("button-dismiss-gap", "", "Size of the gap for dismiss button.")
	rootCmd.Flags().StringP("button-dismiss-no-terminal", "Z", "", "Like --button-dismiss, but does not run the action in a terminal.")
	rootCmd.Flags().String("button-gap", "", "Size of the gap between buttons")
	rootCmd.Flags().String("button-margin-right", "", "Margin from dismiss button to edge.")
	rootCmd.Flags().StringP("button-no-terminal", "B", "", "Like --button, but doesnot run the action in a terminal.")
	rootCmd.Flags().String("button-padding", "", "Padding for the button text.")
	rootCmd.Flags().String("button-text", "", "Button text color.")
	rootCmd.Flags().StringP("config", "c", "", "Path to config file.")
	rootCmd.Flags().BoolP("debug", "d", false, "Enable debugging.")
	rootCmd.Flags().StringP("detailed-button", "L", "", "Set the text of the detail button.")
	rootCmd.Flags().BoolP("detailed-message", "l", false, "Read a detailed message from stdin.")
	rootCmd.Flags().String("details-background", "", "Details background color.")
	rootCmd.Flags().String("details-border-size", "", "Thickness for the details border.")
	rootCmd.Flags().StringP("dismiss-button", "s", "", "Set the dismiss button text.")
	rootCmd.Flags().StringP("edge", "e", "", "Set the edge to use.")
	rootCmd.Flags().StringP("font", "f", "", "Set the font to use.")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message and quit.")
	rootCmd.Flags().StringP("layer", "y", "", "Set the layer to use.")
	rootCmd.Flags().StringP("message", "m", "", "Set the message text.")
	rootCmd.Flags().String("message-padding", "", "Padding for the message.")
	rootCmd.Flags().StringP("output", "o", "", "Set the output to use.")
	rootCmd.Flags().String("text", "", "Text color.")
	rootCmd.Flags().StringP("type", "t", "", "Se, t the message type.")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version number and quit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"background":         color.ActionHexColors(),
		"border":             color.ActionHexColors(),
		"border-bottom":      color.ActionHexColors(),
		"button-background":  color.ActionHexColors(),
		"button-text":        color.ActionHexColors(),
		"config":             carapace.ActionFiles(),
		"details-background": color.ActionHexColors(),
		"edge":               carapace.ActionValues("top", "bottom"),
		"font":               os.ActionFontFamilies(),
		"layer":              carapace.ActionValues("overlay", "top", "bottom", "background"),
		"text":               color.ActionHexColors(),
	})
}
