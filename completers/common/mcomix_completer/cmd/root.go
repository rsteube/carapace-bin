package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcomix",
	Short: "GTK Comic Book Viewer",
	Long:  "https://sourceforge.net/projects/mcomix/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("W", "W", "", "Sets the desired output log level.")
	rootCmd.Flags().BoolP("double-page", "d", false, "Start the application in double page mode.")
	rootCmd.Flags().BoolP("fullscreen", "f", false, "Start the application in fullscreen mode.")
	rootCmd.Flags().Bool("help", false, "Show this help and exit.")
	rootCmd.Flags().BoolP("library", "l", false, "Show the library on startup.")
	rootCmd.Flags().BoolP("manga", "m", false, "Start the application in manga mode.")
	rootCmd.Flags().BoolP("slideshow", "s", false, "Start the application in slideshow mode.")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version number and exit.")
	rootCmd.Flags().BoolP("zoom-best", "b", false, "Start the application with zoom set to best fit mode.")
	rootCmd.Flags().BoolP("zoom-height", "h", false, "Start the application with zoom set to fit height.")
	rootCmd.Flags().BoolP("zoom-width", "w", false, "Start the application with zoom set to fit width.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"W": carapace.ActionValues("all", "debug", "info", "warn", "error").StyleF(style.ForLogLevel),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
