package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lnav",
	Short: "ncurses-based log file viewer",
	Long:  "https://lnav.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "Check configuration and then exit.")
	rootCmd.Flags().BoolS("H", "H", false, "Display the internal help text.")
	rootCmd.Flags().StringS("I", "I", "", "An additional configuration directory.")
	rootCmd.Flags().BoolS("R", "R", false, "Load older rotated log files as well.")
	rootCmd.Flags().BoolS("V", "V", false, "Print version information.")
	rootCmd.Flags().BoolS("a", "a", false, "Load all of the most recent log file types.")
	rootCmd.Flags().StringS("c", "c", "", "Execute a command after the files have been loaded.")
	rootCmd.Flags().StringS("d", "d", "", "Write debug messages to the given file.")
	rootCmd.Flags().StringS("f", "f", "", "Execute the commands in the given file.")
	rootCmd.Flags().BoolS("h", "h", false, "Print this message, then exit.")
	rootCmd.Flags().BoolS("i", "i", false, "Install the given format files and exit.  Pass 'extra'")
	rootCmd.Flags().BoolS("n", "n", false, "Run without the curses UI. (headless mode)")
	rootCmd.Flags().BoolS("q", "q", false, "Do not print the log messages after executing all")
	rootCmd.Flags().BoolS("r", "r", false, "Recursively load files from the given directory hierarchies.")
	rootCmd.Flags().BoolS("t", "t", false, "Prepend timestamps to the lines of data being read in")
	rootCmd.Flags().BoolS("u", "u", false, "Update formats installed from git repositories.")
	rootCmd.Flags().StringS("w", "w", "", "Write the contents of the standard input to this file.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"I": carapace.ActionDirectories(),
		"d": carapace.ActionFiles(),
		"f": carapace.ActionFiles(),
		"w": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
