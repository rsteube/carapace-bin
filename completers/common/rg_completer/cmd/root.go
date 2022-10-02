package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rg_completer/cmd/action"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rg",
	Short: "recursively search current directory for lines matching a pattern",
	Long:  "https://github.com/BurntSushi/ripgrep",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	//TODO descriptions are a mess
	rootCmd.Flags().StringP("after-context", "A", "", "Show NUM lines after each match.")
	rootCmd.Flags().StringP("before-context", "B", "", "Show NUM lines before each match.  ")
	rootCmd.Flags().Bool("binary", false, "Enabling this flag will cause ripgrep to search binary files.")
	rootCmd.Flags().Bool("block-buffered", false, "When enabled, ripgrep will use block buffering")
	rootCmd.Flags().BoolP("byte-offset", "b", false, "Print the 0-based byte offset within the input file before each line of output.")
	rootCmd.Flags().BoolP("case-sensitive", "s", false, "Search case sensitively.")
	rootCmd.Flags().String("color", "", "This flag controls when to use colors.")
	rootCmd.Flags().String("colors", "", "This flag specifies color settings for use in the output.")
	rootCmd.Flags().Bool("column", false, "Show column numbers (1-based).")
	rootCmd.Flags().StringP("context", "C", "", "Show NUM lines before and after each match.")
	rootCmd.Flags().String("context-separator", "", "The string used to separate non-contiguous context lines in the output.")
	rootCmd.Flags().BoolP("count", "c", false, "This flag suppresses normal output and shows the number of lines that match")
	rootCmd.Flags().Bool("count-matches", false, "This flag suppresses normal output and shows the number of individual")
	rootCmd.Flags().Bool("crlf", false, "When enabled, ripgrep will treat CRLF as a line terminator instead")
	rootCmd.Flags().Bool("debug", false, "Show debug messages.")
	rootCmd.Flags().String("dfa-size-limit", "", "The upper size limit of the regex DFA. The default limit is 10M.")
	rootCmd.Flags().StringP("encoding", "E", "", "Specify the text encoding that ripgrep will use on all files searched.")
	rootCmd.Flags().String("engine", "", "Specify which regular expression engine to use.")
	rootCmd.Flags().StringP("file", "f", "", "Search for patterns from the given file, with one pattern per line.")
	rootCmd.Flags().Bool("files", false, "Print each file that would be searched without actually performing the search.")
	rootCmd.Flags().BoolP("files-with-matches", "l", false, "Only print the paths with at least one match.")
	rootCmd.Flags().Bool("files-without-match", false, "Only print the paths that contain zero matches.")
	rootCmd.Flags().BoolP("fixed-strings", "F", false, "Treat the pattern as a literal string instead of a regular expression.")
	rootCmd.Flags().BoolP("follow", "L", false, "When this flag is enabled, ripgrep will follow symbolic links while traversing")
	rootCmd.Flags().StringP("glob", "g", "", "Include or exclude files and directories for searching that match the given")
	rootCmd.Flags().Bool("glob-case-insensitive", false, "Process glob patterns given with the -g/--glob flag case insensitively.")
	rootCmd.Flags().Bool("heading", false, "This flag prints the file path above clusters of matches from each file instead")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information.")
	rootCmd.Flags().Bool("hidden", false, "Search hidden files and directories.")
	rootCmd.Flags().String("iglob", "", "Include or exclude files and directories for searching that match the given")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "When this flag is provided, the given patterns will be searched case")
	rootCmd.Flags().String("ignore-file", "", "Specifies a path to one or more .gitignore format rules files.")
	rootCmd.Flags().Bool("ignore-file-case-insensitive", false, "Process ignore files (.gitignore, .ignore, etc.) case insensitively.")
	rootCmd.Flags().Bool("include-zero", false, "When used with --count or --count-matches, print the number of matches for")
	rootCmd.Flags().BoolP("invert-match", "v", false, "Show lines that do not match the given patterns.")
	rootCmd.Flags().Bool("json", false, "Enable printing results in a JSON Lines format.")
	rootCmd.Flags().Bool("line-buffered", false, "When enabled, ripgrep will use line buffering.")
	rootCmd.Flags().BoolP("line-number", "n", false, "Show line numbers (1-based).")
	rootCmd.Flags().BoolP("line-regexp", "x", false, "Only show matches surrounded by line boundaries.")
	rootCmd.Flags().StringP("max-columns", "M", "", "Don't print lines longer than this limit in bytes.")
	rootCmd.Flags().Bool("max-columns-preview", false, "When the '--max-columns' flag is used, ripgrep will by default completely")
	rootCmd.Flags().StringP("max-count", "m", "", "Limit the number of matching lines per file searched to NUM.")
	rootCmd.Flags().String("max-depth", "", "Limit the depth of directory traversal to NUM levels beyond the paths given.")
	rootCmd.Flags().String("max-filesize", "", "Ignore files larger than NUM in size. This does not apply to directories.")
	rootCmd.Flags().Bool("mmap", false, "Search using memory maps when possible.")
	rootCmd.Flags().BoolP("multiline", "U", false, "Enable matching across multiple lines.")
	rootCmd.Flags().Bool("multiline-dotall", false, "This flag enables \"dot all\" in your regex pattern, which causes '.' to match")
	rootCmd.Flags().Bool("no-config", false, "Never read configuration files.")
	rootCmd.Flags().BoolP("no-filename", "I", false, "Never print the file path with the matched lines.")
	rootCmd.Flags().Bool("no-heading", false, "Don't group matches by each file.")
	rootCmd.Flags().Bool("no-ignore", false, "Don't respect ignore files (.gitignore, .ignore, etc.).")
	rootCmd.Flags().Bool("no-ignore-dot", false, "Don't respect .ignore files.")
	rootCmd.Flags().Bool("no-ignore-exclude", false, "Don't respect ignore files that are manually configured for the repository")
	rootCmd.Flags().Bool("no-ignore-files", false, "When set, any --ignore-file flags, even ones that come after this flag, are")
	rootCmd.Flags().Bool("no-ignore-global", false, "Don't respect ignore files that come from \"global\" sources such as git's")
	rootCmd.Flags().Bool("no-ignore-messages", false, "Suppresses all error messages related to parsing ignore files such as .ignore")
	rootCmd.Flags().Bool("no-ignore-parent", false, "Don't respect ignore files (.gitignore, .ignore, etc.) in parent directories.")
	rootCmd.Flags().Bool("no-ignore-vcs", false, "Don't respect version control ignore files (.gitignore, etc.).")
	rootCmd.Flags().BoolP("no-line-number", "N", false, "Suppress line numbers. This is enabled by default when not searching in a")
	rootCmd.Flags().Bool("no-messages", false, "Suppress all error messages related to opening and reading files.")
	rootCmd.Flags().Bool("no-mmap", false, "Never use memory maps, even when they might be faster.")
	rootCmd.Flags().Bool("no-pcre2-unicode", false, "DEPRECATED. Use --no-unicode instead.")
	rootCmd.Flags().Bool("no-require-git", false, "By default, ripgrep will only respect global gitignore rules, .gitignore rules")
	rootCmd.Flags().Bool("no-unicode", false, "By default, ripgrep will enable \"Unicode mode\" in all of its regexes.")
	rootCmd.Flags().BoolP("null", "0", false, "Whenever a file path is printed, follow it with a NUL byte.")
	rootCmd.Flags().Bool("null-data", false, "Enabling this option causes ripgrep to use NUL as a line terminator instead of")
	rootCmd.Flags().Bool("one-file-system", false, "When enabled, ripgrep will not cross file system boundaries relative to where")
	rootCmd.Flags().BoolP("only-matching", "o", false, "Print only the matched (non-empty) parts of a matching line, with each such")
	rootCmd.Flags().Bool("passthru", false, "Print both matching and non-matching lines.")
	rootCmd.Flags().String("path-separator", "", "Set the path separator to use when printing file paths.")
	rootCmd.Flags().BoolP("pcre2", "P", false, "When this flag is present, ripgrep will use the PCRE2 regex engine instead of")
	rootCmd.Flags().Bool("pcre2-version", false, "When this flag is present, ripgrep will print the version of PCRE2 in use,")
	rootCmd.Flags().String("pre", "", "For each input FILE, search the standard output of COMMAND FILE rather than the")
	rootCmd.Flags().String("pre-glob", "", "This flag works in conjunction with the --pre flag.")
	rootCmd.Flags().BoolP("pretty", "p", false, "This is a convenience alias for '--color always --heading --line-number'.")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print anything to stdout.")
	rootCmd.Flags().String("regex-size-limit", "", "")
	rootCmd.Flags().StringP("regexp", "e", "", "A pattern to search for.")
	rootCmd.Flags().StringP("replace", "r", "", "Replace every match with the text given when printing results.")
	rootCmd.Flags().BoolP("search-zip", "z", false, "Search in compressed files.")
	rootCmd.Flags().BoolP("smart-case", "S", false, "Searches case insensitively if the pattern is all lowercase.")
	rootCmd.Flags().String("sort", "", "This flag enables sorting of results in ascending order.")
	rootCmd.Flags().String("sortr", "", "This flag enables sorting of results in descending order.")
	rootCmd.Flags().Bool("stats", false, "Print aggregate statistics about this ripgrep search.")
	rootCmd.Flags().BoolP("text", "a", false, "Search binary files as if they were text.")
	rootCmd.Flags().StringP("threads", "j", "", "The approximate number of threads to use.")
	rootCmd.Flags().Bool("trim", false, "When set, all ASCII whitespace at the beginning of each line printed will be")
	rootCmd.Flags().StringP("type", "t", "", "Only search files matching TYPE.")
	rootCmd.Flags().String("type-add", "", "Add a new glob for a particular file type.")
	rootCmd.Flags().String("type-clear", "", "Clear the file type globs previously defined for TYPE.")
	rootCmd.Flags().Bool("type-list", false, "Show all supported file types and their corresponding globs.")
	rootCmd.Flags().StringP("type-not", "T", "", "Do not search files matching TYPE.")
	rootCmd.Flags().BoolP("unrestricted", "u", false, "Reduce the level of smart searching.")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.Flags().Bool("vimgrep", false, "Show results with every match on its own line, including line numbers and")
	rootCmd.Flags().BoolP("with-filename", "H", false, "Display the file path for matches.")
	rootCmd.Flags().BoolP("word-regexp", "w", false, "Only show matches surrounded by word boundaries.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValuesDescribed(
			"always", "always use colors",
			"ansi", "always use ANSI colors (even on Windows)",
			"auto", "use colors or not based on stdout, TERM, etc.",
			"never", "never use colors",
		).StyleF(style.ForKeyword),
		"colors": carapace.ActionValuesDescribed(
			"column", "specify coloring for column numbers",
			"line", "specify coloring for line numbers",
			"match", "specify coloring for match text",
			"path", "specify coloring for file names",
		),
		"encoding": carapace.ActionValues("auto", "none", "ASCII", "UTF-8", "UTF-16"),
		"engine": carapace.ActionValuesDescribed(
			"auto", "identical to --auto-hybrid-regex",
			"default", "use default engine",
			"pcre2", "identical to --pcre2",
		),
		"file":        carapace.ActionFiles(),
		"ignore-file": carapace.ActionFiles(),
		"sort": carapace.ActionValuesDescribed(
			"accessed", "sort by last accessed time",
			"created", "sort by creation time",
			"modified", "sort by last modified time",
			"none", "no sorting",
			"path", "sort by file path",
		),
		"sortr": carapace.ActionValuesDescribed(
			"accessed", "sort by last accessed time",
			"created", "sort by creation time",
			"modified", "sort by last modified time",
			"none", "no sorting",
			"path", "sort by file path",
		),
		"type":       action.ActionTypes(),
		"type-add":   action.ActionTypes(),
		"type-clear": action.ActionTypes(),
		"type-not":   action.ActionTypes(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("regexp").Changed || rootCmd.Flag("file").Changed {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
