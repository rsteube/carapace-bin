package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "i3lock",
	Short: "improved screen locker",
	Long:  "https://i3wm.org/i3lock/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("beep", "b", false, "Enable beeping.")
	rootCmd.Flags().StringP("color", "c", "", "Turn the screen into the given color instead of white.")
	rootCmd.Flags().Bool("debug", false, "Enables debug logging.")
	rootCmd.Flags().BoolP("ignore-empty-password", "e", false, "Do not validate empty password.")
	rootCmd.Flags().StringP("image", "i", "", "Display the given PNG image instead of a blank screen.")
	rootCmd.Flags().BoolP("no-unlock-indicator", "u", false, "Disable  the  unlock indicator.")
	rootCmd.Flags().BoolP("nofork", "n", false, "Don't fork after starting.")
	rootCmd.Flags().StringP("pointer", "p", "", "Set pointer.")
	rootCmd.Flags().String("raw", "", "Read the image given by --image as a raw image instead of PNG.")
	rootCmd.Flags().BoolP("show-failed-attempts", "f", false, "Show the number of failed attempts, if any.")
	rootCmd.Flags().BoolP("tiling", "t", false, "If an image is specified (via -i) it will display the image tiled all over the screen")
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of your i3lock")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"image":   carapace.ActionFiles(),
		"pointer": carapace.ActionValues("default", "win"),
	})
}
