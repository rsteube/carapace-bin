package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qmk",
	Short: "CLI wrapper for running QMK commands",
	Long:  "https://github.com/qmk/qmk_cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("color", false, "Enable color in output")
	rootCmd.Flags().String("config-file", "", "The location for the configuration file")
	rootCmd.Flags().String("datetime-fmt", "", "Format string for datetimes")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().Bool("interactive", false, "Force interactive mode even when stdout is not a tty.")
	rootCmd.Flags().String("log-file", "", "File to write log messages to")
	rootCmd.Flags().String("log-file-fmt", "", "Format string for log file.")
	rootCmd.Flags().String("log-file-level", "", "Logging level for log file.")
	rootCmd.Flags().String("log-fmt", "", "Format string for printed log output")
	rootCmd.Flags().Bool("no-color", false, "Disable color in output")
	rootCmd.Flags().Bool("no-unicode", false, "Disable unicode loglevels")
	rootCmd.Flags().Bool("unicode", false, "Enable unicode loglevels")
	rootCmd.Flags().BoolP("verbose", "v", false, "Make the logging more verbose")
	rootCmd.Flags().BoolP("version", "V", false, "Display the version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file":    carapace.ActionFiles(),
		"log-file":       carapace.ActionFiles(),
		"log-file-level": carapace.ActionValues("debug", "info", "warning", "error", "critical").StyleF(style.ForLogLevel),
	})
}
