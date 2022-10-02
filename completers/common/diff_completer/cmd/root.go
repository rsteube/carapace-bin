package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diff",
	Short: "compare files line by line",
	Long:  "https://linux.die.net/man/1/diff",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("GTYPE-group-format", "", "format GTYPE input groups with GFMT")
	rootCmd.Flags().String("LTYPE-line-format", "", "format LTYPE input lines with LFMT")
	rootCmd.Flags().BoolP("brief", "q", false, "report only when files differ")
	rootCmd.Flags().String("color", "", "colorize the output")
	rootCmd.Flags().String("context", "", "output NUM (default 3) lines of copied context")
	rootCmd.Flags().BoolP("ed", "e", false, "output an ed script")
	rootCmd.Flags().StringP("exclude", "x", "", "exclude files that match PAT")
	rootCmd.Flags().StringP("exclude-from", "X", "", "exclude files that match any pattern in FILE")
	rootCmd.Flags().BoolP("expand-tabs", "t", false, "expand tabs to spaces in output")
	rootCmd.Flags().String("from-file", "", "compare FILE1 to all operands")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().String("horizon-lines", "", "keep NUM lines of the common prefix and suffix")
	rootCmd.Flags().StringP("ifdef", "D", "", "output merged file with '#ifdef NAME' diffs")
	rootCmd.Flags().BoolP("ignore-all-space", "w", false, "ignore all white space")
	rootCmd.Flags().BoolP("ignore-blank-lines", "B", false, "ignore changes where lines are all blank")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "ignore case differences in file contents")
	rootCmd.Flags().Bool("ignore-file-name-case", false, "ignore case when comparing file names")
	rootCmd.Flags().StringP("ignore-matching-lines", "I", "", "ignore changes where all lines match RE")
	rootCmd.Flags().BoolP("ignore-space-change", "b", false, "ignore changes in the amount of white space")
	rootCmd.Flags().BoolP("ignore-tab-expansion", "E", false, "ignore changes due to tab expansion")
	rootCmd.Flags().BoolP("ignore-trailing-space", "Z", false, "ignore white space at line end")
	rootCmd.Flags().BoolP("initial-tab", "T", false, "make tabs line up by prepending a tab")
	rootCmd.Flags().String("label", "", "use LABEL instead of file name and timestamp")
	rootCmd.Flags().Bool("left-column", false, "output only the left column of common lines")
	rootCmd.Flags().String("line-format", "", "format all input lines with LFMT")
	rootCmd.Flags().BoolP("minimal", "d", false, "try hard to find a smaller set of changes")
	rootCmd.Flags().BoolP("new-file", "N", false, "treat absent files as empty")
	rootCmd.Flags().Bool("no-dereference", false, "don't follow symbolic links")
	rootCmd.Flags().Bool("no-ignore-file-name-case", false, "consider case when comparing file names")
	rootCmd.Flags().Bool("normal", false, "output a normal diff (the default)")
	rootCmd.Flags().BoolP("paginate", "l", false, "pass output through 'pr' to paginate it")
	rootCmd.Flags().String("palette", "", "the colors to use when --color is active")
	rootCmd.Flags().BoolP("rcs", "n", false, "output an RCS format diff")
	rootCmd.Flags().BoolP("recursive", "r", false, "recursively compare any subdirectories found")
	rootCmd.Flags().BoolP("report-identical-files", "s", false, "report when two files are the same")
	rootCmd.Flags().BoolP("show-c-function", "p", false, "show which C function each change is in")
	rootCmd.Flags().StringP("show-function-line", "F", "", "show the most recent line matching RE")
	rootCmd.Flags().BoolP("side-by-side", "y", false, "output in two columns")
	rootCmd.Flags().Bool("speed-large-files", false, "assume large files and many scattered small changes")
	rootCmd.Flags().StringP("starting-file", "S", "", "start with FILE when comparing directories")
	rootCmd.Flags().Bool("strip-trailing-cr", false, "strip trailing carriage return on input")
	rootCmd.Flags().Bool("suppress-blank-empty", false, "suppress space or tab before empty output lines")
	rootCmd.Flags().Bool("suppress-common-lines", false, "do not output common lines")
	rootCmd.Flags().String("tabsize", "", "tab stops every NUM (default 8) print columns")
	rootCmd.Flags().BoolP("text", "a", false, "treat all files as text")
	rootCmd.Flags().String("to-file", "", "compare all operands to FILE2")
	rootCmd.Flags().Bool("unidirectional-new-file", false, "treat absent first files as empty")
	rootCmd.Flags().String("unified", "", "output NUM (default 3) lines of unified context")
	rootCmd.Flags().BoolP("version", "v", false, "output version information and exit")
	rootCmd.Flags().StringP("width", "W", "", "output at most NUM (default 130) print columns")

	rootCmd.Flag("GTYPE-group-format").NoOptDefVal = " "
	rootCmd.Flag("LTYPE-line-format").NoOptDefVal = " "
	rootCmd.Flag("color").NoOptDefVal = " "
	rootCmd.Flag("context").NoOptDefVal = " "
	rootCmd.Flag("exclude").NoOptDefVal = " "
	rootCmd.Flag("exclude-from").NoOptDefVal = " "
	rootCmd.Flag("from-file").NoOptDefVal = " "
	rootCmd.Flag("horizon-lines").NoOptDefVal = " "
	rootCmd.Flag("ifdef").NoOptDefVal = " "
	rootCmd.Flag("ignore-matching-lines").NoOptDefVal = " "
	rootCmd.Flag("line-format").NoOptDefVal = " "
	rootCmd.Flag("palette").NoOptDefVal = " "
	rootCmd.Flag("show-function-line").NoOptDefVal = " "
	rootCmd.Flag("starting-file").NoOptDefVal = " "
	rootCmd.Flag("to-file").NoOptDefVal = " "
	rootCmd.Flag("unified").NoOptDefVal = " "
	rootCmd.Flag("width").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":        carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"exclude-from": carapace.ActionFiles(),
		"from-file":    carapace.ActionFiles(),
		"to-file":      carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
