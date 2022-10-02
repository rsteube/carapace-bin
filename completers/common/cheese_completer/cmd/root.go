package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cheese",
	Short: "tool to take pictures and videos from your webcam",
	Long:  "https://wiki.gnome.org/Apps/Cheese",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("device", "d", "", "Device to use as a camera")
	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().BoolP("fullscreen", "f", false, "Start in fullscreen mode")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("help-all", false, "Show all help options")
	rootCmd.Flags().Bool("help-gapplication", false, "Show GApplication options")
	rootCmd.Flags().Bool("help-gtk", false, "Show GTK+ Options")
	rootCmd.Flags().BoolP("version", "v", false, "Output version information and exit")
	rootCmd.Flags().BoolP("wide", "w", false, "Start in wide mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"device":  carapace.ActionFiles(),
		"display": os.ActionDisplays(),
	})
}
