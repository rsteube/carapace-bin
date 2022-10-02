package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pcmanfm",
	Short: "A lightweight Gtk+ based file manager for X Window",
	Long:  "https://wiki.lxde.org/en/PCManFM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("class", "", "Program class as used by the window manager")
	rootCmd.Flags().BoolP("daemon-mode", "d", false, "Run PCManFM as a daemon")
	rootCmd.Flags().Bool("desktop", false, "Launch desktop manager")
	rootCmd.Flags().Bool("desktop-off", false, "Turn off desktop manager if it's running")
	rootCmd.Flags().Bool("desktop-pref", false, "Open desktop preference dialog")
	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().BoolP("find-files", "f", false, "Open a Find Files window")
	rootCmd.Flags().Bool("g-fatal-warnings", false, "Make all warnings fatal")
	rootCmd.Flags().String("gtk-module", "", "Load additional GTK+ modules")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("help-all", false, "Show all help options")
	rootCmd.Flags().Bool("help-gtk", false, "Show GTK+ Options")
	rootCmd.Flags().String("name", "", "Program name as used by the window manager")
	rootCmd.Flags().BoolP("new-win", "n", false, "Open new window")
	rootCmd.Flags().Bool("no-desktop", false, "No function. Just to be compatible with nautilus")
	rootCmd.Flags().Bool("one-screen", false, "Use --desktop option only for one screen")
	rootCmd.Flags().StringP("profile", "p", "", "Name of configuration profile")
	rootCmd.Flags().String("role", "", "Window role for usage by window manager")
	rootCmd.Flags().String("screen", "", "X screen to use")
	rootCmd.Flags().StringP("set-wallpaper", "w", "", "Set desktop wallpaper from image FILE")
	rootCmd.Flags().String("show-pref", "", "Open Preferences dialog on the page N")
	rootCmd.Flags().Bool("sync", false, "Make X calls synchronous")
	rootCmd.Flags().String("wallpaper-mode", "", "Set mode of desktop wallpaper")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"display":        os.ActionDisplays(),
		"set-wallpaper":  carapace.ActionFiles(),
		"wallpaper-mode": carapace.ActionValues("color", "stretch", "fit", "crop", "center", "tile", "screen"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
