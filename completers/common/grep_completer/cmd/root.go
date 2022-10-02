package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grep",
	Short: "print lines that match patterns",
	Long:  "https://www.gnu.org/software/grep/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

// ExecuteOverride overrides Use before execution
func ExecuteOverride(use string) error {
	rootCmd.Use = use
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("I", "I", false, "equivalent to --binary-files=without-match")
	//rootCmd.Flags().BoolS("NUM", "NUM", false, "same as --context=NUM")
	rootCmd.Flags().StringP("after-context", "A", "", "print NUM lines of trailing context")
	rootCmd.Flags().BoolP("basic-regexp", "G", false, "PATTERNS are basic regular expressions")
	rootCmd.Flags().StringP("before-context", "B", "", "print NUM lines of leading context")
	rootCmd.Flags().BoolP("binary", "U", false, "do not strip CR characters at EOL (MSDOS/Windows)")
	rootCmd.Flags().String("binary-files", "", "assume that binary files are TYPE;")
	rootCmd.Flags().BoolP("byte-offset", "b", false, "print the byte offset with output lines")
	rootCmd.Flags().String("color", "", "")
	rootCmd.Flags().String("colour", "", "use markers to highlight the matching strings;")
	rootCmd.Flags().StringP("context", "C", "", "print NUM lines of output context")
	rootCmd.Flags().BoolP("count", "c", false, "print only a count of selected lines per FILE")
	rootCmd.Flags().BoolP("dereference-recursive", "R", false, "likewise, but follow all symlinks")
	rootCmd.Flags().StringP("devices", "D", "", "how to handle devices, FIFOs and sockets;")
	rootCmd.Flags().StringP("directories", "d", "", "how to handle directories;")
	rootCmd.Flags().String("exclude", "", "skip files that match GLOB")
	rootCmd.Flags().String("exclude-dir", "", "skip directories that match GLOB")
	rootCmd.Flags().String("exclude-from", "", "skip files that match any file pattern from FILE")
	rootCmd.Flags().BoolP("extended-regexp", "E", false, "PATTERNS are extended regular expressions")
	rootCmd.Flags().StringP("file", "f", "", "take PATTERNS from FILE")
	rootCmd.Flags().BoolP("files-with-matches", "l", false, "print only names of FILEs with selected lines")
	rootCmd.Flags().BoolP("files-without-match", "L", false, "print only names of FILEs with no selected lines")
	rootCmd.Flags().BoolP("fixed-strings", "F", false, "PATTERNS are strings")
	rootCmd.Flags().Bool("help", false, "display this help text and exit")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "ignore case distinctions in patterns and data")
	rootCmd.Flags().String("include", "", "search only files that match GLOB (a file pattern)")
	rootCmd.Flags().BoolP("initial-tab", "T", false, "make tabs line up (if needed)")
	rootCmd.Flags().BoolP("invert-match", "v", false, "select non-matching lines")
	rootCmd.Flags().String("label", "", "use LABEL as the standard input file name prefix")
	rootCmd.Flags().Bool("line-buffered", false, "flush output on every line")
	rootCmd.Flags().BoolP("line-number", "n", false, "print line number with output lines")
	rootCmd.Flags().BoolP("line-regexp", "x", false, "match only whole lines")
	rootCmd.Flags().StringP("max-count", "m", "", "stop after NUM selected lines")
	rootCmd.Flags().BoolP("no-filename", "h", false, "suppress the file name prefix on output")
	rootCmd.Flags().Bool("no-ignore-case", false, "do not ignore case distinctions (default)")
	rootCmd.Flags().BoolP("no-messages", "s", false, "suppress error messages")
	rootCmd.Flags().BoolP("null", "Z", false, "print 0 byte after FILE name")
	rootCmd.Flags().BoolP("null-data", "z", false, "a data line ends in 0 byte, not newline")
	rootCmd.Flags().BoolP("only-matching", "o", false, "show only nonempty parts of lines that match")
	rootCmd.Flags().BoolP("perl-regexp", "P", false, "PATTERNS are Perl regular expressions")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress all normal output")
	rootCmd.Flags().BoolP("recursive", "r", false, "like --directories=recurse")
	rootCmd.Flags().StringP("regexp", "e", "", "use PATTERNS for matching")
	rootCmd.Flags().Bool("silent", false, "suppress all normal output")
	rootCmd.Flags().BoolP("text", "a", false, "equivalent to --binary-files=text")
	rootCmd.Flags().BoolP("version", "V", false, "display version information and exit")
	rootCmd.Flags().BoolP("with-filename", "H", false, "print file name with output lines")
	rootCmd.Flags().BoolP("word-regexp", "w", false, "match only whole words")

	rootCmd.Flag("color").NoOptDefVal = " "
	rootCmd.Flag("colour").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"binary-files": carapace.ActionValues("binary", "test", "without-match"),
		"color":        carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"colour":       carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"devices":      carapace.ActionValues("read", "skip"),
		"directories":  carapace.ActionValues("read", "recurse", "skip"),
		"exclude-from": carapace.ActionFiles(),
		"file":         carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("file").Changed {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
