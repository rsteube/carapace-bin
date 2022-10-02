package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "Print lines matching a pattern",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grepCmd).Standalone()

	grepCmd.Flags().BoolS("H", "H", false, "show filenames")
	grepCmd.Flags().BoolS("I", "I", false, "don't match patterns in binary files")
	grepCmd.Flags().StringP("after-context", "A", "", "show <n> context lines after matches")
	grepCmd.Flags().Bool("all-match", false, "show only matches from files that match all patterns")
	grepCmd.Flags().Bool("and", false, "combine patterns specified with -e")
	grepCmd.Flags().BoolP("basic-regexp", "G", false, "use basic POSIX regular expressions (default)")
	grepCmd.Flags().StringP("before-context", "B", "", "show <n> context lines before matches")
	grepCmd.Flags().Bool("break", false, "print empty line between matches from different files")
	grepCmd.Flags().Bool("cached", false, "search in index instead of in the work tree")
	grepCmd.Flags().String("color", "", "highlight matches")
	grepCmd.Flags().Bool("column", false, "show column number of first match")
	grepCmd.Flags().StringP("context", "C", "", "show <n> context lines before and after matches")
	grepCmd.Flags().BoolP("count", "c", false, "show the number of matches instead of matching lines")
	grepCmd.Flags().StringS("e", "e", "", "match <pattern>")
	grepCmd.Flags().Bool("exclude-standard", false, "ignore files specified via '.gitignore'")
	grepCmd.Flags().Bool("ext-grep", false, "allow calling of grep(1) (ignored by this build)")
	grepCmd.Flags().BoolP("extended-regexp", "E", false, "use extended POSIX regular expressions")
	grepCmd.Flags().StringS("f", "f", "", "read patterns from file")
	grepCmd.Flags().BoolP("files-with-matches", "l", false, "show only filenames instead of matching lines")
	grepCmd.Flags().BoolP("files-without-match", "L", false, "show only the names of files without match")
	grepCmd.Flags().BoolP("fixed-strings", "F", false, "interpret patterns as fixed strings")
	grepCmd.Flags().Bool("full-name", false, "show filenames relative to top directory")
	grepCmd.Flags().BoolS("h", "h", false, "don't show filenames")
	grepCmd.Flags().Bool("heading", false, "show filename only once above matches from same file")
	grepCmd.Flags().BoolP("ignore-case", "i", false, "case insensitive matching")
	grepCmd.Flags().BoolP("invert-match", "v", false, "show non-matching lines")
	grepCmd.Flags().BoolP("line-number", "n", false, "show line numbers")
	grepCmd.Flags().String("max-depth", "", "descend at most <depth> levels")
	grepCmd.Flags().Bool("name-only", false, "synonym for --files-with-matches")
	grepCmd.Flags().Bool("no-index", false, "find in contents not managed by git")
	grepCmd.Flags().Bool("not", false, "")
	grepCmd.Flags().BoolP("null", "z", false, "print NUL after filenames")
	grepCmd.Flags().BoolP("only-matching", "o", false, "show only matching parts of a line")
	grepCmd.Flags().StringP("open-files-in-pager", "O", "", "show matching files in the pager")
	grepCmd.Flags().Bool("or", false, "")
	grepCmd.Flags().BoolP("perl-regexp", "P", false, "use Perl-compatible regular expressions")
	grepCmd.Flags().BoolP("quiet", "q", false, "indicate hit with exit status without output")
	grepCmd.Flags().Bool("recurse-submodules", false, "recursively search in each submodule")
	grepCmd.Flags().BoolP("recursive", "r", false, "search in subdirectories (default)")
	grepCmd.Flags().BoolP("show-function", "p", false, "show a line with the function name before matches")
	grepCmd.Flags().BoolP("text", "a", false, "process binary files as text")
	grepCmd.Flags().Bool("textconv", false, "process binary files with textconv filters")
	grepCmd.Flags().String("threads", "", "use <n> worker threads")
	grepCmd.Flags().Bool("untracked", false, "search in both tracked and untracked files")
	grepCmd.Flags().BoolP("word-regexp", "w", false, "match patterns only at word boundaries")
	rootCmd.AddCommand(grepCmd)

	grepCmd.Flag("color").NoOptDefVal = " "
	grepCmd.Flag("open-files-in-pager").NoOptDefVal = " "

	carapace.Gen(grepCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("always", "never", "auto").StyleF(style.ForKeyword),
		"f":     carapace.ActionFiles(),
		"open-files-in-pager": carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	})

	carapace.Gen(grepCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(grepCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(grepCmd).DashAnyCompletion(
		carapace.ActionFiles(),
	)
}
