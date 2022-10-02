package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calibre",
	Short: "Comprehensive e-book software",
	Long:  "https://calibre-ebook.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("detach", false, "Detach from the controlling terminal, if any (Linux only)")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().Bool("ignore-plugins", false, "Ignore custom plugins")
	rootCmd.Flags().Bool("no-update-check", false, "Do not check for updates")
	rootCmd.Flags().BoolP("shutdown-running-calibre", "s", false, "Cause a running calibre instance, if any, to be shutdown.")
	rootCmd.Flags().Bool("start-in-tray", false, "Start minimized to system tray.")
	rootCmd.Flags().BoolP("verbose", "v", false, "Ignored, do not use. Present only for legacy reasons")
	rootCmd.Flags().Bool("version", false, "show program's version number and exit")
	rootCmd.Flags().String("with-library", "", "Use the library located at the specified path.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"with-library": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
