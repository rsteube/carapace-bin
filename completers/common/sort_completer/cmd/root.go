package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sort",
	Short: "sort lines of text files",
	Long:  "https://linux.die.net/man/1/sort",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("C", "C", "", "--check=silent  like -c, but do not report first bad line")
	rootCmd.Flags().String("batch-size", "", "merge at most NMERGE inputs at once")
	rootCmd.Flags().StringP("buffer-size", "S", "", "use SIZE for main memory buffer")
	rootCmd.Flags().StringP("check", "c", "", "check for sorted input; do not sort")
	rootCmd.Flags().String("compress-program", "", "compress temporaries with PROG")
	rootCmd.Flags().Bool("debug", false, "annotate the part of the line used to sort, and warn about questionable usage to stderr")
	rootCmd.Flags().BoolP("dictionary-order", "d", false, "consider only blanks and alphanumeric characters")
	rootCmd.Flags().StringP("field-separator", "t", "", "use SEP instead of non-blank to blank transition")
	rootCmd.Flags().String("files0-from", "", "read input from the files specified by NUL-terminated names in file F")
	rootCmd.Flags().BoolP("general-numeric-sort", "g", false, "compare according to general numerical value")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("human-numeric-sort", "h", false, "compare human readable numbers (e.g., 2K 1G)")
	rootCmd.Flags().BoolP("ignore-case", "f", false, "fold lower case to upper case characters")
	rootCmd.Flags().BoolP("ignore-leading-blanks", "b", false, "ignore leading blanks")
	rootCmd.Flags().BoolP("ignore-nonprinting", "i", false, "consider only printable characters")
	rootCmd.Flags().StringP("key", "k", "", "sort via a key; KEYDEF gives location and type")
	rootCmd.Flags().BoolP("merge", "m", false, "merge already sorted files; do not sort")
	rootCmd.Flags().BoolP("month-sort", "M", false, "compare (unknown) < 'JAN' < ... < 'DEC'")
	rootCmd.Flags().BoolP("numeric-sort", "n", false, "compare according to string numerical value")
	rootCmd.Flags().StringP("output", "o", "", "write result to FILE instead of standard output")
	rootCmd.Flags().String("parallel", "", "change the number of sorts run concurrently to N")
	rootCmd.Flags().BoolP("random-sort", "R", false, "shuffle, but group identical keys.  See shuf(1)")
	rootCmd.Flags().String("random-source", "", "get random bytes from FILE")
	rootCmd.Flags().BoolP("reverse", "r", false, "reverse the result of comparisons")
	rootCmd.Flags().String("sort", "", "sort according to WORD")
	rootCmd.Flags().BoolP("stable", "s", false, "stabilize sort by disabling last-resort comparison")
	rootCmd.Flags().StringP("temporary-directory", "T", "", "use DIR for temporaries, not $TMPDIR or /tmp;")
	rootCmd.Flags().BoolP("unique", "u", false, "with -c, check for strict ordering; without -c, output only the first of an equal run")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("version-sort", "V", false, "natural sort of (version) numbers within text")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"check":               carapace.ActionValues("diagnose-first", "quiet", "silent"),
		"compress-program":    carapace.ActionValues("bzip2", "gzip", "lzop", "xz"),
		"files0-from":         carapace.ActionFiles(),
		"random-source":       carapace.ActionFiles(),
		"sort":                carapace.ActionValues("general-numeric", "human-numeric", "month", "numeric", "random", "version"),
		"temporary-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
