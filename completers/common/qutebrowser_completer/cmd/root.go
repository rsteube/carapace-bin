package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qutebrowser",
	Short: "a keyboard-driven, vim-like browser based on PyQt5",
	Long:  "https://qutebrowser.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("backend", "", "Which backend to use")
	rootCmd.Flags().StringP("basedir", "B", "", "Base directory for all storage")
	rootCmd.Flags().StringP("config-py", "C", "", "Path to config.py")
	rootCmd.Flags().BoolP("debug", "d", false, "Turn on debugging options")
	rootCmd.Flags().StringP("debug-flag", "D", "", "Pass name of debugging feature to be turned on")
	rootCmd.Flags().Bool("enable-webengine-inspector", false, "Enable the web inspector")
	rootCmd.Flags().Bool("force-color", false, "Force colored logging")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().Bool("json-logging", false, "Output log lines in JSON format (one object per line)")
	rootCmd.Flags().String("logfilter", "", "Comma-separated list of things to be logged")
	rootCmd.Flags().StringP("loglevel", "l", "", "Override the configured console loglevel")
	rootCmd.Flags().String("loglines", "", "How many lines of the debug log to keep in RAM")
	rootCmd.Flags().Bool("no-err-windows", false, "Don't show any error windows")
	rootCmd.Flags().Bool("nocolor", false, "Turn off colored logging")
	rootCmd.Flags().Bool("nowindow", false, "Don't show the main window")
	rootCmd.Flags().BoolP("override-restore", "R", false, "Don't restore a session even if one would be restored")
	rootCmd.Flags().String("qt-flag", "", "Pass an argument to Qt as flag")
	rootCmd.Flags().StringP("restore", "r", "", "Restore a named session")
	rootCmd.Flags().String("target", "", "How URLs should be opened")
	rootCmd.Flags().BoolP("temp-basedir", "T", false, "Use a temporary basedir")
	rootCmd.Flags().BoolP("version", "V", false, "Show version and quit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"backend":   carapace.ActionValues("webkit", "webengine"),
		"basedir":   carapace.ActionDirectories(),
		"config-py": carapace.ActionFiles(),
		"loglevel":  carapace.ActionValues("critical", "error", "warning", "info", "debug", "vdebug").StyleF(style.ForLogLevel),
		"target":    carapace.ActionValues("auto", "tab", "tab-bg", "tab-silent", "tab-bg-silent", "window"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
