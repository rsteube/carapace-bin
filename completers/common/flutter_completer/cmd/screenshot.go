package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var screenshotCmd = &cobra.Command{
	Use:   "screenshot",
	Short: "Take a screenshot from a connected device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(screenshotCmd).Standalone()

	screenshotCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach.")
	screenshotCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	screenshotCmd.Flags().String("observatory-url", "", "The Observatory URL to which to connect.")
	screenshotCmd.Flags().StringP("out", "o", "", "Location to write the screenshot.")
	screenshotCmd.Flags().String("type", "", "The type of screenshot to retrieve.")
	rootCmd.AddCommand(screenshotCmd)

	carapace.Gen(screenshotCmd).FlagCompletion(carapace.ActionMap{
		"out": carapace.ActionFiles(),
		"type": carapace.ActionValuesDescribed(
			"device", "Delegate to the device's native screenshot capabilities.",
			"rasterizer", "Render the Flutter app using the rasterizer.",
			"skia", "Render the Flutter app as a Skia picture.",
		),
	})
}
