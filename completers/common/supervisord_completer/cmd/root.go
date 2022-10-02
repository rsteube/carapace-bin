package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "supervisord",
	Short: "run a set of applications as daemons",
	Long:  "http://supervisord.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("childlogdir", "q", "", "the log directory for child process logs")
	rootCmd.Flags().StringP("configuration", "c", "", "configuration file path ")
	rootCmd.Flags().StringP("directory", "d", "", "directory to chdir to when daemonized")
	rootCmd.Flags().BoolP("help", "h", false, "print this usage message and exit")
	rootCmd.Flags().StringP("identifier", "i", "", "identifier used for this instance of supervisord")
	rootCmd.Flags().StringP("logfile", "l", "", "use FILENAME as logfile path")
	rootCmd.Flags().StringP("logfile_backups", "z", "", "number of backups to keep when max bytes reached")
	rootCmd.Flags().StringP("logfile_maxbytes", "y", "", "use BYTES to limit the max size of logfile")
	rootCmd.Flags().StringP("loglevel", "e", "", "use LEVEL as log level ")
	rootCmd.Flags().StringP("minfds", "a", "", "the minimum number of file descriptors for start success")
	rootCmd.Flags().String("minprocs", "", "the minimum number of processes available for start success")
	rootCmd.Flags().BoolP("nocleanup", "k", false, "prevent the process from performing cleanup ")
	rootCmd.Flags().BoolP("nodaemon", "n", false, "run in the foreground ")
	rootCmd.Flags().StringP("pidfile", "j", "", "write a pid file for the daemon process to FILENAME")
	rootCmd.Flags().String("profile_options", "", "run supervisord under profiler and output results")
	rootCmd.Flags().BoolP("silent", "s", false, "no logs to stdout ")
	rootCmd.Flags().BoolP("strip_ansi", "t", false, "strip ansi escape codes from process output")
	rootCmd.Flags().StringP("umask", "m", "", "use this umask for daemon subprocess ")
	rootCmd.Flags().StringP("user", "u", "", "run supervisord as this user ")
	rootCmd.Flags().BoolP("version", "v", false, "print supervisord version number and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"childlogdir":   carapace.ActionDirectories(),
		"configuration": carapace.ActionFiles(),
		"directory":     carapace.ActionDirectories(),
		"logfile":       carapace.ActionFiles(),
		"loglevel":      carapace.ActionValues("trace", "debug", "info", "warn", "error", "critical").StyleF(style.ForLogLevel),
		"pidfile":       carapace.ActionFiles(),
		"profile_options": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("cumulative", "calls", "callers").Invoke(c).Filter(c.Parts).ToA()
		}),
		"umask": fs.ActionFileModesNumeric(),
		"user":  os.ActionUsers(),
	})
}
