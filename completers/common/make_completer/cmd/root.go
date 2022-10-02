package cmd

import (
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/make"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "make",
	Short: "GNU make utility to maintain groups of programs",
	Long:  "https://linux.die.net/man/1/make",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("always-make", "B", false, "Unconditionally make all targets.")
	rootCmd.Flags().String("assume-new", "", "Consider FILE to be infinitely new.")
	rootCmd.Flags().String("assume-old", "", "Consider FILE to be very old and don't remake it.")
	rootCmd.Flags().BoolP("check-symlink-times", "L", false, "Use the latest mtime between symlinks and target.")
	rootCmd.Flags().BoolS("d", "d", false, "Print lots of debugging information.")
	rootCmd.Flags().String("debug", "", "Print various types of debugging information.")
	rootCmd.Flags().StringP("directory", "C", "", "Change to DIRECTORY before doing anything.")
	rootCmd.Flags().Bool("dry-run", false, "Don't actually run any recipe; just print them.")
	rootCmd.Flags().BoolP("environment-overrides", "e", false, "Environment variables override makefiles.")
	rootCmd.Flags().StringP("eval", "E", "", "Evaluate STRING as a makefile statement.")
	rootCmd.Flags().StringP("file", "f", "", "Read FILE as a makefile.")
	rootCmd.Flags().BoolP("help", "h", false, "Print this message and exit.")
	rootCmd.Flags().BoolP("ignore-errors", "i", false, "Ignore errors from recipes.")
	rootCmd.Flags().StringP("include-dir", "I", "", "Search DIRECTORY for included makefiles.")
	rootCmd.Flags().StringP("jobs", "j", "", "Allow N jobs at once; infinite jobs with no arg.")
	rootCmd.Flags().BoolP("just-print", "n", false, "Don't actually run any recipe; just print them.")
	rootCmd.Flags().BoolP("keep-going", "k", false, "Keep going when some targets can't be made.")
	rootCmd.Flags().StringP("load-average", "l", "", "Don't start multiple jobs unless load is below N.")
	rootCmd.Flags().String("makefile", "", "Read FILE as a makefile.")
	rootCmd.Flags().String("max-load", "", "Don't start multiple jobs unless load is below N.")
	rootCmd.Flags().String("new-file", "", "Consider FILE to be infinitely new.")
	rootCmd.Flags().BoolP("no-builtin-rules", "r", false, "Disable the built-in implicit rules.")
	rootCmd.Flags().BoolP("no-builtin-variables", "R", false, "Disable the built-in variable settings.")
	rootCmd.Flags().BoolP("no-keep-going", "S", false, "Turns off -k.")
	rootCmd.Flags().Bool("no-print-directory", false, "Turn off -w, even if it was turned on implicitly.")
	rootCmd.Flags().Bool("no-silent", false, "Echo recipes (disable --silent mode).")
	rootCmd.Flags().StringP("old-file", "o", "", "Consider FILE to be very old and don't remake it.")
	rootCmd.Flags().StringP("output-sync", "O", "", "Synchronize output of parallel jobs by TYPE.")
	rootCmd.Flags().BoolP("print-data-base", "p", false, "Print make's internal database.")
	rootCmd.Flags().BoolP("print-directory", "w", false, "Print the current directory.")
	rootCmd.Flags().BoolP("question", "q", false, "Run no recipe; exit status says if up to date.")
	rootCmd.Flags().Bool("quiet", false, "Don't echo recipes.")
	rootCmd.Flags().Bool("recon", false, "Don't actually run any recipe; just print them.")
	rootCmd.Flags().BoolP("silent", "s", false, "Don't echo recipes.")
	rootCmd.Flags().Bool("stop", false, "Turns off -k.")
	rootCmd.Flags().BoolP("touch", "t", false, "Touch targets instead of remaking them.")
	rootCmd.Flags().Bool("trace", false, "Print tracing information.")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number of make and exit.")
	rootCmd.Flags().Bool("warn-undefined-variables", false, "Warn when an undefined variable is referenced.")
	rootCmd.Flags().StringP("what-if", "W", "", "Consider FILE to be infinitely new.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"assume-new":  carapace.ActionFiles(),
		"assume-old":  carapace.ActionFiles(),
		"directory":   carapace.ActionDirectories(),
		"file":        carapace.ActionFiles(),
		"include-dir": carapace.ActionDirectories(),
		"new-file":    carapace.ActionFiles(),
		"old-file":    carapace.ActionFiles(),
		"what-if":     carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			file := "Makefile"
			if rootCmd.Flag("directory").Changed {
				file = fmt.Sprintf("%v/Makefile", rootCmd.Flag("directory").Value.String())
			}
			if rootCmd.Flag("file").Changed {
				file = rootCmd.Flag("file").Value.String()
			}
			return make.ActionTargets(file).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
