package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gdb",
	Short: "This is the GNU debugger",
	Long:  "https://www.sourceware.org/gdb/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("args", false, "Arguments after executable-file are passed to inferior.")
	rootCmd.Flags().StringS("b", "b", "", "Set serial port baud rate used for remote debugging.")
	rootCmd.Flags().Bool("batch", false, "Exit after processing options.")
	rootCmd.Flags().Bool("batch-silent", false, "Like --batch, but suppress all gdb stdout output.")
	rootCmd.Flags().String("cd", "", "Change current directory to DIR.")
	rootCmd.Flags().StringP("command", "x", "", "Execute GDB commands from FILE.")
	rootCmd.Flags().Bool("configuration", false, "Print details about GDB configuration and then exit.")
	rootCmd.Flags().String("core", "", "Analyze the core dump COREFILE.")
	rootCmd.Flags().StringP("data-directory", "D", "", "Set GDB's data-directory to DIR.")
	rootCmd.Flags().Bool("dbx", false, "DBX compatibility mode.")
	rootCmd.Flags().String("directory", "", "Search for source files in DIR.")
	rootCmd.Flags().String("eval-command", "", "Execute a single GDB command.")
	rootCmd.Flags().String("exec", "", "Use EXECFILE as the executable.")
	rootCmd.Flags().Bool("fullname", false, "Output information used by emacs-GDB interface.")
	rootCmd.Flags().Bool("help", false, "Print this message and then exit.")
	rootCmd.Flags().String("init-command", "", "Like -x but execute commands before loading inferior.")
	rootCmd.Flags().String("init-eval-command", "", "Like -ex but before loading inferior.")
	rootCmd.Flags().String("interpreter", "", "Select a specific interpreter / user interface.")
	rootCmd.Flags().StringS("l", "l", "", "Set timeout in seconds for remote debugging.")
	rootCmd.Flags().Bool("nh", false, "Do not read ~/.gdbinit.")
	rootCmd.Flags().Bool("nw", false, "Do not use the GUI interface.")
	rootCmd.Flags().Bool("nx", false, "Do not read any .gdbinit files in any directory.")
	rootCmd.Flags().String("pid", "", "Attach to running process PID.")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print version number on startup.")
	rootCmd.Flags().Bool("readnever", false, "Do not read symbol files.")
	rootCmd.Flags().Bool("readnow", false, "Fully read symbol files on first access.")
	rootCmd.Flags().Bool("return-child-result", false, "GDB exit code will be the child's exit code.")
	rootCmd.Flags().String("se", "", "Use FILE as symbol file and executable file.")
	rootCmd.Flags().Bool("silent", false, "Do not print version number on startup.")
	rootCmd.Flags().String("symbols", "", "Read symbols from SYMFILE.")
	rootCmd.Flags().String("tty", "", "Use TTY for input/output by the program being debugged.")
	rootCmd.Flags().Bool("tui", false, "Use a terminal user interface.")
	rootCmd.Flags().Bool("version", false, "Print version information and then exit.")
	rootCmd.Flags().BoolS("w", "w", false, "Use the GUI interface.")
	rootCmd.Flags().Bool("write", false, "Set writing into executable and core files.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cd":             carapace.ActionDirectories(),
		"command":        carapace.ActionFiles(),
		"core":           carapace.ActionFiles(),
		"data-directory": carapace.ActionDirectories(),
		"directory":      carapace.ActionDirectories(),
		"exec":           carapace.ActionFiles(),
		"interpreter": carapace.ActionValuesDescribed(
			"console", "The traditional console or command-line interpreter.",
			"mi", "The newest GDB/MI interface (currently mi2).",
			"mi2", "The current GDB/MI interface.",
			"mi1", "The GDB/MI interface included in GDB 5.1, 5.2, and 5.3.",
		),
		"pid":     ps.ActionProcessIds(),
		"se":      carapace.ActionFiles(),
		"symbols": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
