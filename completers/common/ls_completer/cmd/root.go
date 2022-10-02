package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ls",
	Short: "list directory contents",
	Long:  "https://linux.die.net/man/1/ls",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("1", "1", false, "list one file per line.  Avoid '\n' with -q or -b")
	rootCmd.Flags().BoolS("C", "C", false, "list entries by columns")
	rootCmd.Flags().BoolS("S", "S", false, "sort by file size, largest first")
	rootCmd.Flags().BoolS("U", "U", false, "do not sort; list entries in directory order")
	rootCmd.Flags().BoolS("X", "X", false, "sort alphabetically by entry extension")
	rootCmd.Flags().BoolP("all", "a", false, "do not ignore entries starting with .")
	rootCmd.Flags().BoolP("almost-all", "A", false, "do not list implied . and ..")
	rootCmd.Flags().Bool("author", false, "with -l, print the author of each file")
	rootCmd.Flags().String("block-size", "", "with -l, scale sizes by SIZE when printing them;")
	rootCmd.Flags().BoolS("c", "c", false, "with -lt: sort by, and show, ctime (time of last")
	rootCmd.Flags().BoolP("classify", "F", false, "append indicator (one of */=>@|) to entries")
	rootCmd.Flags().String("color", "", "colorize the output; WHEN can be 'always' (default")
	rootCmd.Flags().BoolP("context", "Z", false, "print any security context of each file")
	rootCmd.Flags().BoolP("dereference", "L", false, "when showing file information for a symbolic")
	rootCmd.Flags().BoolP("dereference-command-line", "H", false, "follow symbolic links listed on the command line")
	rootCmd.Flags().Bool("dereference-command-line-symlink-to-dir", false, "follow each command line symbolic link that points to a directory")
	rootCmd.Flags().BoolP("directory", "d", false, "list directories themselves, not their contents")
	rootCmd.Flags().BoolP("dired", "D", false, "generate output designed for Emacs' dired mode")
	rootCmd.Flags().BoolP("escape", "b", false, "print C-style escapes for nongraphic characters")
	rootCmd.Flags().BoolS("f", "f", false, "do not sort, enable -aU, disable -ls --color")
	rootCmd.Flags().Bool("file-type", false, "likewise, except do not append '*'")
	rootCmd.Flags().String("format", "", "across -x, commas -m, horizontal -x, long -l,")
	rootCmd.Flags().Bool("full-time", false, "like -l --time-style=full-iso")
	rootCmd.Flags().BoolS("g", "g", false, "like -l, but do not list owner")
	rootCmd.Flags().Bool("group-directories-first", false, "group directories first")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().String("hide", "", "do not list implied entries matching shell PATTERN")
	rootCmd.Flags().BoolP("hide-control-chars", "q", false, "print ? instead of nongraphic characters")
	rootCmd.Flags().BoolP("human-readable", "h", false, "with -l and -s, print sizes like 1K 234M 2G etc.")
	rootCmd.Flags().String("hyperlink", "", "hyperlink file names; WHEN can be 'always'")
	rootCmd.Flags().StringP("ignore", "I", "", "do not list implied entries matching shell PATTERN")
	rootCmd.Flags().BoolP("ignore-backups", "B", false, "do not list implied entries ending with ~")
	rootCmd.Flags().String("indicator-style", "p", "append indicator with style WORD to entry names:")
	rootCmd.Flags().BoolP("inode", "i", false, "print the index number of each file")
	rootCmd.Flags().BoolP("kibibytes", "k", false, "default to 1024-byte blocks for disk usage;")
	rootCmd.Flags().BoolS("l", "l", false, "use a long listing format")
	rootCmd.Flags().BoolP("literal", "N", false, "print entry names without quoting")
	rootCmd.Flags().BoolS("m", "m", false, "fill width with a comma separated list of entries")
	rootCmd.Flags().BoolP("no-group", "G", false, "in a long listing, don't print group names")
	rootCmd.Flags().BoolP("numeric-uid-gid", "n", false, "like -l, but list numeric user and group IDs")
	rootCmd.Flags().BoolS("o", "o", false, "like -l, but do not list group information")
	rootCmd.Flags().BoolP("quote-name", "Q", false, "enclose entry names in double quotes")
	rootCmd.Flags().String("quoting-style", "", "use quoting style WORD for entry names:")
	rootCmd.Flags().BoolP("recursive", "R", false, "list subdirectories recursively")
	rootCmd.Flags().BoolP("reverse", "r", false, "reverse order while sorting")
	rootCmd.Flags().Bool("show-control-chars", false, "show nongraphic characters as-is (the default,")
	rootCmd.Flags().Bool("si", false, "likewise, but use powers of 1000 not 1024")
	rootCmd.Flags().BoolP("size", "s", false, "print the allocated size of each file, in blocks")
	rootCmd.Flags().String("sort", "", "sort by WORD instead of name: none (-U), size (-S),")
	rootCmd.Flags().BoolS("t", "t", false, "sort by time, newest first; see --time")
	rootCmd.Flags().StringP("tabsize", "T", "", "assume tab stops at each COLS instead of 8")
	rootCmd.Flags().String("time", "", "change the default of using modification times;")
	rootCmd.Flags().String("time-style", "", "time/date format with -l; see TIME_STYLE below")
	rootCmd.Flags().BoolS("u", "u", false, "with -lt: sort by, and show, access time;")
	rootCmd.Flags().BoolS("v", "v", false, "natural sort of (version) numbers within text")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().StringP("width", "w", "", "set output width to COLS.  0 means no limit")
	rootCmd.Flags().BoolS("x", "x", false, "list entries by lines instead of by columns")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":           carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"format":          carapace.ActionValues("across", "horizontal", "single-column", "vertical", "commas", "long", "verbose"),
		"hyperlink":       carapace.ActionValues("always", "auto", "none").StyleF(style.ForKeyword),
		"indicator-style": carapace.ActionValues("classify", "file-type", "none", "slash"),
		"quoting-style":   carapace.ActionValues("c", "clocale", "escape", "literal", "locale", "shell", "shell-always"),
		"sort":            carapace.ActionValues("extension", "none", "size", "time", "version"),
		"time":            carapace.ActionValues("access", "atime", "ctime", "status", "use"),
		"time-style":      carapace.ActionValues("full-iso", "iso", "locale", "long-iso"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
