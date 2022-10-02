package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().StringS("G", "G", "", "Look for differences whose patch text contains added/removed lines that match <regex>.")
	logCmd.Flags().StringS("L", "L", "", "Trace the evolution of the line range given by <start>,<end>, or by the function name regex <funcname>, within the <file>.")
	logCmd.Flags().StringS("O", "O", "", "Control the order in which files appear in the output.")
	logCmd.Flags().BoolS("R", "R", false, "Swap two inputs")
	logCmd.Flags().StringS("S", "S", "", "Look for differences that change the number of occurrences of the specified string.")
	logCmd.Flags().String("abbrev", "", "Show the shortest prefix that is at least <n> hexdigits long that uniquely refers")
	logCmd.Flags().Bool("abbrev-commit", false, "Instead of showing the full 40-byte hexadecimal commit object name, show a prefix that names the object uniquely.")
	logCmd.Flags().String("after", "", "Show commits more recent than a specific date.")
	logCmd.Flags().Bool("all", false, "Pretend as if all the refs in refs/, along with HEAD, are listed on the command line as <commit>.")
	logCmd.Flags().Bool("all-match", false, "Limit the commits output to ones that match all given --grep, instead of ones that match at least one.")
	logCmd.Flags().Bool("alternate-refs", false, "Pretend as if all objects mentioned as ref tips of alternate repositories were listed on the command line.")
	logCmd.Flags().Bool("ancestry-path", false, "When given a range of commits to display, only display commits that exist directly on the ancestry chain.")
	logCmd.Flags().String("anchored", "", "Generate a diff using the \"anchored diff\" algorithm.")
	logCmd.Flags().String("author", "", "Limit the commits output to ones with author/committer header lines that match the specified pattern")
	logCmd.Flags().Bool("author-date-order", false, "Show no parents before all of its children are shown, but otherwise show commits in the author timestamp order.")
	logCmd.Flags().Bool("basic-regexp", false, "Consider the limiting patterns to be basic regular expressions; this is the default.")
	logCmd.Flags().String("before", "", "Show commits older than a specific date.")
	logCmd.Flags().Bool("binary", false, "In addition to --full-index, output a binary diff that can be applied with git-apply. Implies --patch.")
	logCmd.Flags().Bool("bisect", false, "Pretend as if the bad bisection ref refs/bisect/bad was listed and as if it was followed by --not and the good bisection refs refs/bisect/good-* on the command line.")
	logCmd.Flags().Bool("boundary", false, "Output excluded boundary commits. Boundary commits are prefixed with -.")
	logCmd.Flags().String("branches", "", "Pretend as if all the refs in refs/heads are listed on the command line as <commit>.")
	logCmd.Flags().StringP("break-rewrites", "B", "", "Break complete rewrite changes into pairs of delete and create.")
	logCmd.Flags().Bool("check", false, "Warn if changes introduce conflict markers or whitespace errors.")
	logCmd.Flags().Bool("cherry", false, "A synonym for --right-only --cherry-mark --no-merges")
	logCmd.Flags().Bool("cherry-mark", false, "Like --cherry-pick (see below) but mark equivalent commits with = rather than omitting them, and inequivalent ones with +.")
	logCmd.Flags().Bool("cherry-pick", false, "Omit any commit that introduces the same change as another commit on the “other side” when the set of commits are limited with symmetric difference.")
	logCmd.Flags().Bool("children", false, "Print also the children of the commit.")
	logCmd.Flags().String("color", "", "Show colored diff.")
	logCmd.Flags().String("color-moved", "", "Moved lines of code are colored differently.")
	logCmd.Flags().String("color-moved-ws", "", "This configures how whitespace is ignored when performing the move detection for --color-moved.")
	logCmd.Flags().String("color-words", "", "Equivalent to --word-diff=color plus (if a regex was specified) --word-diff-regex=<regex>.")
	logCmd.Flags().Bool("combined-all-paths", false, "This flag causes combined diffs (used for merge commits) to list the name of the file from all parents.")
	logCmd.Flags().String("committer", "", "Limit the commits output to ones with author/committer header lines that match the specified pattern")
	logCmd.Flags().Bool("compact-summary", false, "Output a condensed summary of extended header information.")
	logCmd.Flags().Bool("cumulative", false, "Synonym for --dirstat=cumulative")
	logCmd.Flags().String("date", "", "Only takes effect for dates shown in human-readable format, such as when using --pretty.")
	logCmd.Flags().Bool("date-order", false, "Show no parents before all of its children are shown, but otherwise show commits in the commit timestamp order.")
	logCmd.Flags().String("decorate", "", "Print out the ref names of any commits that are shown.")
	logCmd.Flags().String("decorate-refs", "", "Decorate refs matching pattern.")
	logCmd.Flags().String("decorate-refs-exclude", "", "Do not decorete refs matching pattern.")
	logCmd.Flags().String("diff-algorithm", "", "Choose a diff algorithm.")
	logCmd.Flags().String("diff-filter", "", "Select only files that are of given change.")
	logCmd.Flags().String("diff-merges", "", "Specify diff format to be used for merge commits.")
	logCmd.Flags().StringP("dirstat", "X", "", "Output the distribution of relative amount of changes for each sub-directory.")
	logCmd.Flags().String("dirstat-by-file", "", "Synonym for --dirstat=files,param1,param2...")
	logCmd.Flags().Bool("do-walk", false, "Overrides a previous --no-walk.")
	logCmd.Flags().String("dst-prefix", "", "Show the given destination prefix instead of \"b/\".")
	logCmd.Flags().String("encoding", "", "Commit objects record the character encoding used for the log message in their encoding header.")
	logCmd.Flags().String("exclude", "", "Do not include refs matching <glob-pattern> that the next --all, --branches, --tags, --remotes, or --glob would otherwise consider.")
	logCmd.Flags().Bool("exclude-first-parent-only", false, "When finding commits to exclude (with a ^), follow only the first parent commit upon seeing a merge commit.")
	logCmd.Flags().String("expand-tabs", "", "Perform a tab expansion")
	logCmd.Flags().Bool("ext-diff", false, "Allow an external diff helper to be executed.")
	logCmd.Flags().BoolP("extended-regexp", "E", false, "Consider the limiting patterns to be extended regular expressions instead of the default basic regular expressions.")
	logCmd.Flags().StringP("find-copies", "C", "", "Detect copies as well as renames.")
	logCmd.Flags().Bool("find-copies-harder", false, "For performance reasons, by default, -C option finds copies only if the original file of the copy was modified in the same changeset.")
	logCmd.Flags().String("find-object", "", "Look for differences that change the number of occurrences of the specified object.")
	logCmd.Flags().Bool("first-parent", false, "When finding commits to include, follow only the first parent commit upon seeing a merge commit.")
	logCmd.Flags().BoolP("fixed-strings", "F", false, "Consider the limiting patterns to be fixed strings (don’t interpret pattern as a regular expression).")
	logCmd.Flags().Bool("follow", false, "Continue listing the history of a file beyond renames (works only for a single file).")
	logCmd.Flags().String("format", "", "Pretty-print the contents of the commit logs in a given format")
	logCmd.Flags().Bool("full-diff", false, "With this, the full diff is shown for commits that touch the specified paths")
	logCmd.Flags().Bool("full-history", false, "Same as the default mode, but does not prune some history.")
	logCmd.Flags().Bool("full-index", false, "Instead of the first handful of characters, show the full pre- and post-image blob object names on the \"index\" line when generating patch format output.")
	logCmd.Flags().BoolP("function-context", "W", false, "Show whole function as context lines for each change.")
	logCmd.Flags().String("glob", "", "Pretend as if all the refs matching shell glob <glob-pattern> are listed on the command line as <commit>.")
	logCmd.Flags().Bool("graph", false, "Draw a text-based graphical representation of the commit history on the left hand side of the output.")
	logCmd.Flags().String("grep", "", "Limit the commits output to ones with log message that matches the specified pattern")
	logCmd.Flags().String("grep-reflog", "", "Limit the commits output to ones with reflog entries that match the specified pattern")
	logCmd.Flags().Bool("histogram", false, "Generate a diff using the \"histogram diff\" algorithm.")
	logCmd.Flags().BoolP("ignore-all-space", "w", false, "Ignore whitespace when comparing lines.")
	logCmd.Flags().Bool("ignore-blank-lines", false, "Ignore changes whose lines are all blank.")
	logCmd.Flags().Bool("ignore-cr-at-eol", false, "Ignore carriage-return at the end of line when doing a comparison.")
	logCmd.Flags().StringP("ignore-matching-lines", "I", "", "Ignore changes whose all lines match <regex>.")
	logCmd.Flags().Bool("ignore-missing", false, "Upon seeing an invalid object name in the input, pretend as if the bad input was not given.")
	logCmd.Flags().Bool("ignore-space-at-eol", false, "Ignore changes in whitespace at EOL.")
	logCmd.Flags().BoolP("ignore-space-change", "b", false, "Ignore changes in amount of whitespace.")
	logCmd.Flags().String("ignore-submodules", "", "Ignore changes to submodules in the diff generation.")
	logCmd.Flags().Bool("indent-heuristic", false, "Enable the heuristic that shifts diff hunk boundaries to make patches easier to read. This is the default.")
	logCmd.Flags().String("inter-hunk-context", "", "Show the context between diff hunks, up to the specified number of lines, thereby fusing hunks that are close to each other.")
	logCmd.Flags().Bool("invert-grep", false, "Limit the commits output to ones with log message that do not match the pattern specified with --grep=<pattern>.")
	logCmd.Flags().BoolP("irreversible-delete", "D", false, "Omit the preimage for deletes, i.e. print only the header but not the diff between the preimage and /dev/null.")
	logCmd.Flags().Bool("ita-invisible-in-index", false, ". This option makes the entry appear as a new file in \"git diff\" and non-existent in \"git diff --cached\".")
	logCmd.Flags().StringS("l", "l", "", "This option prevents the exhaustive portion of rename/copy detection from running if the number of source/destination files involved exceeds the specified number.")
	logCmd.Flags().Bool("left-only", false, "List only commits on the respective side of a symmetric difference.")
	logCmd.Flags().Bool("left-right", false, "Mark which side of a symmetric difference a commit is reachable from.")
	logCmd.Flags().String("line-prefix", "", "Prepend an additional prefix to every line of output.")
	logCmd.Flags().Bool("log-size", false, "Include a line “log size <number>” in the output for each commit.")
	logCmd.Flags().Bool("mailmap", false, "Use mailmap file to map author and committer names and email addresses to canonical real names and email addresses.")
	logCmd.Flags().StringP("max-count", "n", "", "Limit the number of commits to output.")
	logCmd.Flags().String("max-parents", "", "Show only commits which have at least (or at most) that many parent commits.")
	logCmd.Flags().Bool("merge", false, "After a failed merge, show refs that touch files having a conflict and don’t exist on all heads to merge.")
	logCmd.Flags().Bool("merges", false, "Print only merge commits. This is exactly the same as --min-parents=2.")
	logCmd.Flags().String("min-parents", "", "Show only commits which have at least (or at most) that many parent commits.")
	logCmd.Flags().Bool("minimal", false, "Spend extra time to make sure the smallest possible diff is produced.")
	logCmd.Flags().Bool("name-only", false, "Show only names of changed files.")
	logCmd.Flags().Bool("name-status", false, "Show only names and status of changed files.")
	logCmd.Flags().Bool("no-abbrev-commit", false, "Show the full 40-byte hexadecimal commit object name.")
	logCmd.Flags().Bool("no-color", false, "Turn off colored diff.")
	logCmd.Flags().Bool("no-color-moved", false, "Turn off move detection.")
	logCmd.Flags().Bool("no-color-moved-ws", false, "Do not ignore whitespace when performing move detection.")
	logCmd.Flags().Bool("no-decorate", false, "Print out the ref names of any commits that are shown.")
	logCmd.Flags().Bool("no-diff-merges", false, "Specify diff format to be used for merge commits.")
	logCmd.Flags().Bool("no-expand-tabs", false, "Do not perform a tab expansion")
	logCmd.Flags().Bool("no-ext-diff", false, "Disallow external diff drivers.")
	logCmd.Flags().Bool("no-indent-heuristic", false, "Disable the indent heuristic.")
	logCmd.Flags().Bool("no-mailmap", false, "Do not use mailmap file to map author and committer names and email addresses to canonical real names and email addresses.")
	logCmd.Flags().Bool("no-max-parents", false, "Show only commits which have at least (or at most) that many parent commits.")
	logCmd.Flags().Bool("no-merges", false, "Do not print commits with more than one parent. This is exactly the same as --max-parents=1.")
	logCmd.Flags().Bool("no-min-parents", false, "Show only commits which have at least (or at most) that many parent commits.")
	logCmd.Flags().Bool("no-notes", false, "Do not show notes.")
	logCmd.Flags().BoolP("no-patch", "s", false, "Suppress diff output.")
	logCmd.Flags().Bool("no-prefix", false, "Do not show any source or destination prefix.")
	logCmd.Flags().Bool("no-relative", false, "When run from a subdirectory of the project, it can be told to exclude changes outside the directory and show pathnames relative to it with this option.")
	logCmd.Flags().Bool("no-rename-empty", false, "Whether to use empty blobs as rename source.")
	logCmd.Flags().Bool("no-renames", false, "Turn off rename detection, even when the configuration file gives the default to do so.")
	logCmd.Flags().Bool("no-textconv", false, "Disallow external text conversion filters to be run when comparing binary files.")
	logCmd.Flags().Bool("no-use-mailmap", false, "Do not use mailmap file to map author and committer names and email addresses to canonical real names and email addresses.")
	logCmd.Flags().String("no-walk", "", "Only show the given commits, but do not traverse their ancestors.")
	logCmd.Flags().Bool("not", false, "Reverses the meaning of the ^ prefix (or lack thereof) for all following revision specifiers, up to the next --not.")
	logCmd.Flags().String("notes", "", "Show the notes that annotate the commit")
	logCmd.Flags().Bool("numstat", false, "Similar to --stat, but shows number of added and deleted lines in decimal notation and pathname without abbreviation.")
	logCmd.Flags().Bool("oneline", false, "This is a shorthand for \"--pretty=oneline --abbrev-commit\" used together.")
	logCmd.Flags().String("output", "", "Output to a specific file instead of stdout.")
	logCmd.Flags().String("output-indicator-context", "", "Specify the character used to indicate context lines in the generated patch.")
	logCmd.Flags().String("output-indicator-new", "", "Specify the character used to indicate new lines in the generated patch.")
	logCmd.Flags().String("output-indicator-old", "", "Specify the character used to indicate old lines in the generated patch.")
	logCmd.Flags().Bool("parents", false, "Print also the parents of the commit.")
	logCmd.Flags().BoolP("patch", "p", false, "Generate patch.")
	logCmd.Flags().Bool("patch-with-raw", false, "Synonym for -p --raw.")
	logCmd.Flags().Bool("patch-with-stat", false, "Synonym for -p --stat.")
	logCmd.Flags().Bool("patience", false, "Generate a diff using the \"patience diff\" algorithm.")
	logCmd.Flags().BoolP("perl-regexp", "P", false, "Consider the limiting patterns to be Perl-compatible regular expressions.")
	logCmd.Flags().Bool("pickaxe-all", false, "When -S or -G finds a change, show all the changes in that changeset, not just the files that contain the change in <string>.")
	logCmd.Flags().Bool("pickaxe-regex", false, "Treat the <string> given to -S as an extended POSIX regular expression to match.")
	logCmd.Flags().String("pretty", "", "Pretty-print the contents of the commit logs in a given format")
	logCmd.Flags().Bool("raw", false, "For each commit, show a summary of changes using the raw diff format.")
	logCmd.Flags().Bool("reflog", false, "Pretend as if all objects mentioned by reflogs are listed on the command line as <commit>.")
	logCmd.Flags().BoolP("regexp-ignore-case", "i", false, "Match the regular expression limiting patterns without regard to letter case.")
	logCmd.Flags().String("relative", "", "When run from a subdirectory of the project, it can be told to exclude changes outside the directory and show pathnames relative to it with this option.")
	logCmd.Flags().Bool("relative-date", false, "Synonym for --date=relative.")
	logCmd.Flags().String("remotes", "", "Pretend as if all the refs in refs/remotes are listed on the command line as <commit>.")
	logCmd.Flags().Bool("remove-empty", false, "Stop when a given path disappears from the tree.")
	logCmd.Flags().Bool("rename-empty", false, "Whether to use empty blobs as rename source.")
	logCmd.Flags().Bool("reverse", false, "Output the commits chosen to be shown (see Commit Limiting section above) in reverse order. Cannot be combined with --walk-reflogs.")
	logCmd.Flags().Bool("right-only", false, "List only commits on the respective side of a symmetric difference.")
	logCmd.Flags().String("rotate-to", "", "Discard the files before the named <file> from the output (i.e.  skip to), or move them to the end of the output (i.e.  rotate to).")
	logCmd.Flags().Bool("shortstat", false, "Output only the last line of the --stat format containing total number of modified files, as well as number of added and deleted lines.")
	logCmd.Flags().String("show-linear-break", "", "When --graph is not used, all history branches are flattened which can make it hard to see that the two consecutive commits do not belong to a linear branch.")
	logCmd.Flags().Bool("show-pulls", false, "Include all commits from the default mode, but also any merge commits that are not TREESAME to the first parent but are TREESAME to a later parent.")
	logCmd.Flags().Bool("show-signature", false, "Check the validity of a signed commit object by passing the signature to gpg --verify and show the output.")
	logCmd.Flags().Bool("simplify-by-decoration", false, "Commits that are referred by some branch or tag are selected.")
	logCmd.Flags().Bool("simplify-merges", false, "First, build a history graph in the same way that --full-history with parent rewriting does.")
	logCmd.Flags().String("since", "", "Show commits more recent than a specific date.")
	logCmd.Flags().Bool("single-worktree", false, "This option forces them to examine the current working tree only.")
	logCmd.Flags().String("skip", "", "Skip number commits before starting to show the commit output.")
	logCmd.Flags().String("skip-to", "", "Discard the files before the named <file> from the output (i.e.  skip to), or move them to the end of the output (i.e.  rotate to).")
	logCmd.Flags().Bool("source", false, "Print out the ref name given on the command line by which each commit was reached.")
	logCmd.Flags().Bool("sparse", false, "All commits that are walked are included.")
	logCmd.Flags().String("src-prefix", "", "Show the given source prefix instead of \"a/\".")
	logCmd.Flags().String("stat", "", "Generate a diffstat.")
	logCmd.Flags().Bool("stdin", false, "In addition to the <commit> listed on the command line, read them from the standard input.")
	logCmd.Flags().String("submodule", "", "Specify how differences in submodules are shown.")
	logCmd.Flags().Bool("summary", false, "Output a condensed summary of extended header information such as creations, renames and mode changes.")
	logCmd.Flags().BoolS("t", "t", false, "Show the tree objects in the diff output.")
	logCmd.Flags().String("tags", "", "Pretend as if all the refs in refs/tags are listed on the command line as <commit>.")
	logCmd.Flags().BoolP("text", "a", false, "Treat all files as text.")
	logCmd.Flags().Bool("textconv", false, "Allow external text conversion filters to be run when comparing binary files.")
	logCmd.Flags().Bool("topo-order", false, "Show no parents before all of its children are shown, and avoid showing commits on multiple lines of history intermixed.")
	logCmd.Flags().StringP("unified", "U", "", "Generate diffs with <n> lines of context instead of the usual three.")
	logCmd.Flags().String("until", "", "Show commits older than a specific date.")
	logCmd.Flags().Bool("use-mailmap", false, "Use mailmap file to map author and committer names and email addresses to canonical real names and email addresses.")
	logCmd.Flags().BoolP("walk-reflogs", "g", false, "Instead of walking the commit ancestry chain, walk reflog entries from the most recent one to older ones.")
	logCmd.Flags().String("word-diff", "", "Show a word diff, using the <mode> to delimit changed words.")
	logCmd.Flags().String("word-diff-regex", "", "Use <regex> to decide what a word is, instead of considering runs of non-whitespace to be a word.")
	logCmd.Flags().String("ws-error-highlight", "", "Highlight whitespace errors in the context, old or new lines of the diff.")
	logCmd.Flags().BoolS("z", "z", false, "Separate the commits with NULs instead of with new newlines.")
	rootCmd.AddCommand(logCmd)

	logCmd.Flag("color").NoOptDefVal = " "
	logCmd.Flag("color-moved").NoOptDefVal = " "
	logCmd.Flag("color-moved-ws").NoOptDefVal = " "
	logCmd.Flag("decorate").NoOptDefVal = "short"
	logCmd.Flag("diff-filter").NoOptDefVal = " "
	logCmd.Flag("format").NoOptDefVal = " "
	logCmd.Flag("ignore-submodules").NoOptDefVal = " "
	logCmd.Flag("no-walk").NoOptDefVal = " "
	logCmd.Flag("pretty").NoOptDefVal = " "
	logCmd.Flag("word-diff").NoOptDefVal = " "
	logCmd.Flag("ws-error-highlight").NoOptDefVal = " "

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"after":  time.ActionDate(),
		"author": git.ActionAuthors(),
		"before": time.ActionDate(),
		"color":  carapace.ActionValues("always", "never", "auto").StyleF(style.ForKeyword),
		"color-moved": carapace.ActionValuesDescribed(
			"no", "Moved lines are not highlighted",
			"default", "Is a synonym for zebra",
			"plain", "Any line that is added in one location and was removed in another location will be colored with color.diff.newMoved",
			"blocks", "Blocks of moved text of at least 20 alphanumeric characters are detected greedily",
			"zebra", "Blocks of moved text are detected as in blocks mode",
			"dimmed-zebra", "Similar to zebra, but additional dimming of uninteresting parts of moved code is performed",
		),
		"color-moved-ws": carapace.ActionValuesDescribed(
			"no", "Do not ignore whitespace when performing move detection",
			"ignore-space-at-eol", "Ignore changes in whitespace at EOL",
			"ignore-space-change", "Ignore changes in amount of whitespace",
			"ignore-all-space", "Ignore whitespace when comparing lines",
			"allow-indentation-change", "Initially ignore any whitespace in the move detection",
		),
		"committer": git.ActionCommitters(),
		"date": carapace.ActionValuesDescribed(
			"relative", "shows dates relative to the current time",
			"local", "is an alias for --date=default-local",
			"iso", "shows timestamps in a ISO 8601-like format",
			"iso8601", "shows timestamps in a ISO 8601-like format",
			"iso-strict", "shows timestamps in strict ISO 8601 format",
			"iso8601-strict", "shows timestamps in strict ISO 8601 format",
			"rfc", "shows timestamps in RFC 2822 format",
			"rfc2822", "shows timestamps in RFC 2822 format",
			"short", "shows only the date, but not the time, in YYYY-MM-DD format",
			"raw", "shows the date as seconds since the epoch (1970-01-01 00:00:00 UTC)",
			"human", "shows the timezone if the timezone does not match the current time-zone",
			"unix", "shows the date as a Unix epoch timestamp (seconds since 1970)",
			"format:", "feeds the format to your system strftime",
			"default", "is the default format",
		),
		"decorate": carapace.ActionValuesDescribed(
			"short", "do not print ref prefixes",
			"full", "print ref prefixes",
			"auto", "short format when output to terminal",
			"no", "no decoration",
		),
		"diff-algorithm": carapace.ActionValuesDescribed(
			"default", "The basic greedy diff algorithm.",
			"myers", "The basic greedy diff algorithm.",
			"minimal", "Spend extra time to make sure the smallest possible diff is produced.",
			"patience", "Use \"patience diff\" algorithm when generating patches.",
			"histogram", "This algorithm extends the patience algorithm to \"support low-occurrence common elements\".",
		),
		"diff-filter": carapace.ActionMultiParts("", func(c carapace.Context) carapace.Action {
			return carapace.ActionValuesDescribed(
				"A", "Added",
				"C", "Copied",
				"D", "Deleted",
				"M", "Modified",
				"R", "Renamed",
				"T", "Changed",
				"U", "Unmerged",
				"X", "Unknown",
				"B", "Broken",
				"*", "All-or-none",
			).Invoke(c).Filter(c.Parts).ToA()
		}),
		"diff-merges": carapace.ActionValuesDescribed(
			"off", "Disable output of diffs for merge commits.",
			"none", "Disable output of diffs for merge commits.",
			"on", "This option makes diff output for merge commits to be shown in the default format.",
			"m", "This option makes diff output for merge commits to be shown in the default format.",
			"first-parent", "This option makes merge commits show the full diff with respect to the first parent only.",
			"1", "This option makes merge commits show the full diff with respect to the first parent only.",
			"separate", "This makes merge commits show the full diff with respect to each of the parents.",
			"remerge", "With this option, two-parent merge commits are remerged to create a temporary tree object",
			"r", "With this option, two-parent merge commits are remerged to create a temporary tree object",
			"combined", "With this option, diff output for a merge commit shows the differences from each of the parents to the merge result simultaneously.",
			"c", "With this option, diff output for a merge commit shows the differences from each of the parents to the merge result simultaneously.",
			"dense-combined", "With this option the output produced by --diff-merges=combined is further compressed.",
			"cc", "With this option the output produced by --diff-merges=combined is further compressed.",
		),
		"format":            carapace.ActionValues("oneline", "short", "medium", "full", "fuller", "reference", "email", "raw", "format:", "tformat:"),
		"ignore-submodules": carapace.ActionValues("none", "untracked", "dirty", "all"),
		"no-walk":           carapace.ActionValues("sorted", "unsorted"),
		"output":            carapace.ActionFiles(),
		"pretty":            carapace.ActionValues("oneline", "short", "medium", "full", "fuller", "reference", "email", "raw", "format:", "tformat:"),
		"since":             time.ActionDate(),
		"until":             time.ActionDate(),
		"word-diff": carapace.ActionValuesDescribed(
			"color", "Highlight changed words using only colors",
			"plain", "Show words as [-removed-] and {+added+}",
			"porcelain", "Use a special line-based format intended for script consumption",
			"none", "Disable word diff again",
		),
		"ws-error-highlight": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("context", "old", "new", "none", "all", "default").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(logCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionMultiParts("...", func(c carapace.Context) carapace.Action {
					if len(c.Parts) < 2 {
						return git.ActionRefs(git.RefOption{}.Default())
					} else {
						return carapace.ActionValues()
					}
				})
			}
		}),
	)
}
