package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/color"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swaylock",
	Short: "Screen locker for Wayland",
	Long:  "https://github.com/swaywm/swaylock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("bs-hl-color", "", "Sets the color of backspace highlight segments.")
	rootCmd.Flags().String("caps-lock-bs-hl-color", "", "Sets the color of backspace highlight segments when Caps Lock is active.")
	rootCmd.Flags().String("caps-lock-key-hl-color", "", "Sets the color of the key press highlight segments when Caps Lock is active.")
	rootCmd.Flags().StringP("color", "c", "", "Turn the screen into the given color instead of white.")
	rootCmd.Flags().StringP("config", "C", "", "Path to the config file.")
	rootCmd.Flags().BoolP("daemonize", "f", false, "Detach from the controlling terminal after locking.")
	rootCmd.Flags().BoolP("debug", "d", false, "Enable debugging output.")
	rootCmd.Flags().BoolP("disable-caps-lock-text", "L", false, "Disable the Caps Lock text.")
	rootCmd.Flags().String("font", "", "Sets the font of the text.")
	rootCmd.Flags().String("font-size", "", "Sets a fixed font size for the indicator text.")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message and quit.")
	rootCmd.Flags().BoolP("hide-keyboard-layout", "K", false, "Hide the current xkb layout while typing.")
	rootCmd.Flags().BoolP("ignore-empty-password", "e", false, "When an empty password is provided, do not validate it.")
	rootCmd.Flags().StringP("image", "i", "", "Display the given image, optionally only on the given output.")
	rootCmd.Flags().BoolP("indicator-caps-lock", "l", false, "Show the current Caps Lock state also on the indicator.")
	rootCmd.Flags().Bool("indicator-idle-visible", false, "Sets the indicator to show even if idle.")
	rootCmd.Flags().String("indicator-radius", "", "Sets the indicator radius.")
	rootCmd.Flags().String("indicator-thickness", "", "Sets the indicator thickness.")
	rootCmd.Flags().String("indicator-x-position", "", "Sets the horizontal position of the indicator.")
	rootCmd.Flags().String("indicator-y-position", "", "Sets the vertical position of the indicator.")
	rootCmd.Flags().String("inside-caps-lock-color", "", "Sets the color of the inside of the indicator when Caps Lock is active.")
	rootCmd.Flags().String("inside-clear-color", "", "Sets the color of the inside of the indicator when cleared.")
	rootCmd.Flags().String("inside-color", "", "Sets the color of the inside of the indicator.")
	rootCmd.Flags().String("inside-ver-color", "", "Sets the color of the inside of the indicator when verifying.")
	rootCmd.Flags().String("inside-wrong-color", "", "Sets the color of the inside of the indicator when invalid.")
	rootCmd.Flags().String("key-hl-color", "", "Sets the color of the key press highlight segments.")
	rootCmd.Flags().String("layout-bg-color", "", "Sets the background color of the box containing the layout text.")
	rootCmd.Flags().String("layout-border-color", "", "Sets the color of the border of the box containing the layout text.")
	rootCmd.Flags().String("layout-text-color", "", "Sets the color of the layout text.")
	rootCmd.Flags().String("line-caps-lock-color", "", "Sets the color of the line between the inside and ring when Caps Lock is active.")
	rootCmd.Flags().String("line-clear-color", "", "Sets the color of the line between the inside and ring when cleared.")
	rootCmd.Flags().String("line-color", "", "Sets the color of the line between the inside and ring.")
	rootCmd.Flags().BoolP("line-uses-inside", "n", false, "Use the inside color for the line between the inside and ring.")
	rootCmd.Flags().BoolP("line-uses-ring", "r", false, "Use the ring color for the line between the inside and ring.")
	rootCmd.Flags().String("line-ver-color", "", "Sets the color of the line between the inside and ring when verifying.")
	rootCmd.Flags().String("line-wrong-color", "", "Sets the color of the line between the inside and ring when invalid.")
	rootCmd.Flags().BoolP("no-unlock-indicator", "u", false, "Disable the unlock indicator.")
	rootCmd.Flags().String("ring-caps-lock-color", "", "Sets the color of the ring of the indicator when Caps Lock is active.")
	rootCmd.Flags().String("ring-clear-color", "", "Sets the color of the ring of the indicator when cleared.")
	rootCmd.Flags().String("ring-color", "", "Sets the color of the ring of the indicator.")
	rootCmd.Flags().String("ring-ver-color", "", "Sets the color of the ring of the indicator when verifying.")
	rootCmd.Flags().String("ring-wrong-color", "", "Sets the color of the ring of the indicator when invalid.")
	rootCmd.Flags().StringP("scaling", "s", "", "Image scaling mode: stretch, fill, fit, center, tile, solid_color.")
	rootCmd.Flags().String("separator-color", "", "Sets the color of the lines that separate highlight segments.")
	rootCmd.Flags().BoolP("show-failed-attempts", "F", false, "Show current count of failed authentication attempts.")
	rootCmd.Flags().BoolP("show-keyboard-layout", "k", false, "Display the current xkb layout while typing.")
	rootCmd.Flags().String("text-caps-lock-color", "", "Sets the color of the text when Caps Lock is active.")
	rootCmd.Flags().String("text-clear-color", "", "Sets the color of the text when cleared.")
	rootCmd.Flags().String("text-color", "", "Sets the color of the text.")
	rootCmd.Flags().String("text-ver-color", "", "Sets the color of the text when verifying.")
	rootCmd.Flags().String("text-wrong-color", "", "Sets the color of the text when invalid.")
	rootCmd.Flags().BoolP("tiling", "t", false, "Same as --scaling=tile.")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version number and quit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bs-hl-color":            color.ActionHexColors(),
		"caps-lock-bs-hl-color":  color.ActionHexColors(),
		"caps-lock-key-hl-color": color.ActionHexColors(),
		"color":                  color.ActionHexColors(),
		"config":                 carapace.ActionFiles(),
		"font":                   os.ActionFontFamilies(),
		"inside-caps-lock-color": color.ActionHexColors(),
		"inside-clear-color":     color.ActionHexColors(),
		"inside-color":           color.ActionHexColors(),
		"inside-ver-color":       color.ActionHexColors(),
		"inside-wrong-color":     color.ActionHexColors(),
		"key-hl-color":           color.ActionHexColors(),
		"layout-bg-color":        color.ActionHexColors(),
		"layout-border-color":    color.ActionHexColors(),
		"layout-text-color":      color.ActionHexColors(),
		"line-caps-lock-color":   color.ActionHexColors(),
		"line-clear-color":       color.ActionHexColors(),
		"line-color":             color.ActionHexColors(),
		"line-ver-color":         color.ActionHexColors(),
		"line-wrong-color":       color.ActionHexColors(),
		"ring-caps-lock-color":   color.ActionHexColors(),
		"ring-clear-color":       color.ActionHexColors(),
		"ring-color":             color.ActionHexColors(),
		"ring-ver-color":         color.ActionHexColors(),
		"ring-wrong-color":       color.ActionHexColors(),
		"scaling":                carapace.ActionValues("stretch", "fill", "fit", "center", "tile", "solid_color"),
		"separator-color":        color.ActionHexColors(),
		"text-caps-lock-color":   color.ActionHexColors(),
		"text-clear-color":       color.ActionHexColors(),
		"text-color":             color.ActionHexColors(),
		"text-ver-color":         color.ActionHexColors(),
		"text-wrong-color":       color.ActionHexColors(),
	})
}
