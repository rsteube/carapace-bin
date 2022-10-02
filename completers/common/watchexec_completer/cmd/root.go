package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "watchexec",
	Short: "Execute commands when watched files change",
	Long:  "https://github.com/watchexec/watchexec",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("changes-only", false, "Only print path change information. Overridden by --verbose")
	rootCmd.Flags().BoolP("clear", "c", false, "Clear screen before executing command")
	rootCmd.Flags().StringP("debounce", "d", "", "Set the timeout between detected change and command execution, defaults to 100ms")
	rootCmd.Flags().StringP("exts", "e", "", "Comma-separated list of file extensions to watch")
	rootCmd.Flags().StringP("filter", "f", "", "Ignore all modifications except those matching the pattern")
	rootCmd.Flags().String("force-poll", "", "Force polling mode (interval in milliseconds)")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().StringP("ignore", "i", "", "Ignore modifications to paths matching the pattern")
	rootCmd.Flags().Bool("no-default-ignore", false, "Skip auto-ignoring of commonly ignored globs")
	rootCmd.Flags().Bool("no-environment", false, "Do not set WATCHEXEC_*_PATH environment variables for the command")
	rootCmd.Flags().Bool("no-ignore", false, "Skip auto-loading of ignore files (.gitignore, .ignore, etc.) for filtering")
	rootCmd.Flags().Bool("no-meta", false, "Ignore metadata changes")
	rootCmd.Flags().Bool("no-process-group", false, "Do not use a process group when running the command")
	rootCmd.Flags().BoolP("no-shell", "n", false, "Do not wrap command in a shell. Deprecated: use --shell=none instead.")
	rootCmd.Flags().Bool("no-vcs-ignore", false, "Skip auto-loading of .gitignore files for filtering")
	rootCmd.Flags().BoolP("notify", "N", false, "Send a desktop notification when watchexec notices a change")
	rootCmd.Flags().String("on-busy-update", "", "Select the behaviour to use when receiving events while the command is running.")
	rootCmd.Flags().BoolP("postpone", "p", false, "Wait until first change to execute command")
	rootCmd.Flags().BoolP("restart", "r", false, "Restart the process if it's still running.")
	rootCmd.Flags().String("shell", "", "Use a different shell, or `none`. E.g. --shell=bash")
	rootCmd.Flags().StringP("signal", "s", "", "Send signal to process upon changes, e.g. SIGHUP")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print debugging messages to stderr")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.Flags().StringP("watch", "w", "", "Watch a specific file or directory")
	rootCmd.Flags().BoolP("watch-when-idle", "W", false, "Deprecated alias for --on-busy-update=do-nothing, which will become the")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"exts": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return fs.ActionFilenameExtensions().Invoke(c).Filter(c.Parts).ToA()
		}),
		"on-busy-update": carapace.ActionValues("do-nothing", "queue", "restart", "signal"),
		"shell":          os.ActionShells(),
		"signal":         ps.ActionKillSignals(),
		"watch":          carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
