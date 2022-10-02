package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show changes between commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	addDiffFlags(diffCmd)

	carapace.Gen(diffCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)
}

func addDiffFlags(cmd *cobra.Command) {
	cmd.Flags().BoolS("0", "0", false, "Omit diff output for unmerged entries")
	cmd.Flags().StringS("O", "O", "", "Control the order in which files appear in the output.")
	cmd.Flags().BoolS("R", "R", false, "Swap two inputs")
	cmd.Flags().StringS("S", "S", "", "Look for differences that change the number of occurrences of the specified string")
	cmd.Flags().String("abbrev", "", "show only a partial prefix")
	cmd.Flags().String("anchored", "", "Generate a diff using the \"anchored diff\" algorithm.")
	cmd.Flags().BoolP("base", "1", false, "compare with base")
	cmd.Flags().Bool("binary", false, "output a binary diff")
	cmd.Flags().StringP("break-rewrites", "B", "", "Break complete rewrite changes into pairs of delete and create.")
	cmd.Flags().Bool("check", false, "Warn if changes introduce conflict markers or whitespace errors")
	cmd.Flags().String("color", "", "Show colored diff.")
	cmd.Flags().String("color-moved", "", "Moved lines of code are colored differently.")
	cmd.Flags().String("color-moved-ws", "", "This configures how whitespace is ignored when performing the move detection")
	cmd.Flags().String("color-words", "", "Equivalent to --word-diff=color plus (if a regex was specified)")
	cmd.Flags().Bool("compact-summary", false, "Output a condensed summary of extended header information")
	cmd.Flags().Bool("cumulative", false, "Synonym for --dirstat=cumulative")
	cmd.Flags().String("diff-algorithm", "", "")
	cmd.Flags().String("diff-filter", "", "filter files")
	cmd.Flags().StringP("dirstat", "X", "", "Output the distribution of relative amount of changes for each sub-directory")
	cmd.Flags().String("dirstat-by-file", "", "Synonym for --dirstat=files,param1,param2...")
	cmd.Flags().String("dst-prefix", "", "Show the given destination prefix instead of \"b/\".")
	cmd.Flags().Bool("exit-code", false, "Make the program exit with codes similar to diff(1).")
	cmd.Flags().Bool("ext-diff", false, "Allow an external diff helper to be executed.")
	cmd.Flags().StringP("find-copies", "C", "", "Detect copies as well as renames")
	cmd.Flags().Bool("find-copies-harder", false, "inspect unmodified files as candidates for the source of copy")
	cmd.Flags().String("find-object", "", "Look for differences that change the number of occurrences of the specified object.")
	cmd.Flags().StringP("find-renames", "M", "", "Detect renames")
	cmd.Flags().Bool("full-index", false, "show the full pre- and post-image blob object names")
	cmd.Flags().BoolP("function-context", "W", false, "Show whole surrounding functions of changes.")
	cmd.Flags().Bool("histogram", false, "Generate a diff using the \"histogram diff\" algorithm.")
	cmd.Flags().BoolP("ignore-all-space", "w", false, "Ignore whitespace when comparing lines.")
	cmd.Flags().Bool("ignore-blank-lines", false, "Ignore changes whose lines are all blank.")
	cmd.Flags().Bool("ignore-cr-at-eol", false, "Ignore carriage-return at the end of line when doing a comparison.")
	cmd.Flags().Bool("ignore-space-at-eol", false, "Ignore changes in whitespace at EOL.")
	cmd.Flags().BoolP("ignore-space-change", "b", false, "Ignore changes in amount of whitespace.")
	cmd.Flags().String("ignore-submodules", "", "Ignore changes to submodules in the diff generation.")
	cmd.Flags().Bool("indent-heuristic", false, "Enable the heuristic that shifts diff hunk boundaries to make patches easier to read.")
	cmd.Flags().String("inter-hunk-context", "", "Show the context between diff hunks")
	cmd.Flags().BoolP("irreversible-delete", "D", false, "Omit the preimage for deletes")
	cmd.Flags().Bool("ita-invisible-in-index", false, "this option makes the entry appear as a new file")
	cmd.Flags().StringS("l", "l", "", "prevent rename/copy detection from running if the number of rename/copy targets exceeds the specified number")
	cmd.Flags().String("line-prefix", "", "Prepend an additional prefix to every line of output.")
	cmd.Flags().Bool("minimal", false, "Spend extra time to make sure the smallest possible diff is produced.")
	cmd.Flags().Bool("name-only", false, "Show only names of changed files.")
	cmd.Flags().Bool("name-status", false, "Show only names and status of changed files.")
	cmd.Flags().Bool("no-color", false, "Turn off colored diff.")
	cmd.Flags().Bool("no-color-moved", false, "Turn off move detection.")
	cmd.Flags().Bool("no-color-moved-ws", false, "Do not ignore whitespace when performing move detection")
	cmd.Flags().Bool("no-ext-diff", false, "Disallow external diff drivers.")
	cmd.Flags().Bool("no-indent-heuristic", false, "Disable the indent heuristic.")
	cmd.Flags().BoolP("no-patch", "s", false, "Suppress diff output.")
	cmd.Flags().Bool("no-prefix", false, "Do not show any source or destination prefix.")
	cmd.Flags().Bool("no-rename-empty", false, "Whether to use empty blobs as rename source.")
	cmd.Flags().Bool("no-renames", false, "Turn off rename detection")
	cmd.Flags().Bool("numstat", false, "Similar to --stat, but shows number of added and deleted lines in decimal notation")
	cmd.Flags().BoolP("ours", "2", false, "compare with our branch")
	cmd.Flags().String("output", "", "Output to a specific file instead of stdout.")
	cmd.Flags().String("output-indicator-context", "", "Specify the character used to indicate context lines in the generated patch.")
	cmd.Flags().String("output-indicator-new", "", "Specify the character used to indicate new lines in the generated patch.")
	cmd.Flags().String("output-indicator-old", "", "Specify the character used to indicate old lines in the generated patch.")
	cmd.Flags().BoolP("patch", "p", false, "Generate patch")
	cmd.Flags().Bool("patch-with-raw", false, "Synonym for -p --raw.")
	cmd.Flags().Bool("patch-with-stat", false, "Synonym for -p --stat.")
	cmd.Flags().Bool("patience", false, "Generate a diff using the \"patience diff\" algorithm.")
	cmd.Flags().Bool("pickaxe-all", false, "When -S or -G finds a change, show all the changes in that changeset")
	cmd.Flags().Bool("pickaxe-regex", false, "Treat the <string> given to -S as an extended POSIX regular expression to match.")
	cmd.Flags().Bool("quiet", false, "Disable all output of the program.")
	cmd.Flags().Bool("raw", false, "Generate the diff in raw format.")
	cmd.Flags().String("relative,", "", "exclude changes outside the directory")
	cmd.Flags().Bool("rename-empty", false, "Whether to use empty blobs as rename source.")
	cmd.Flags().Bool("shortstat", false, "Output only the last line of the --stat format")
	cmd.Flags().String("src-prefix", "", "Show the given source prefix instead of \"a/\".")
	cmd.Flags().String("stat", "", "")
	cmd.Flags().String("submodule", "", "Specify how differences in submodules are shown.")
	cmd.Flags().Bool("summary", false, "Output a condensed summary of extended header information")
	cmd.Flags().BoolP("text", "a", false, "Treat all files as text.")
	cmd.Flags().String("textconv,", "", "Allow (or disallow) external text conversion filters to be run when comparing binary files")
	cmd.Flags().BoolP("theirs", "3", false, "compare with their branch")
	cmd.Flags().BoolS("u", "u", false, "Generate patch")
	cmd.Flags().StringP("unified", "U", "", "Generate diffs with <n> lines of context instead of the usual three.")
	cmd.Flags().String("word-diff", "", "Show a word diff, using the <mode> to delimit changed words.")
	cmd.Flags().String("word-diff-regex", "", "Use <regex> to decide what a word is")
	cmd.Flags().String("ws-error-highlight", "", "Highlight whitespace errors in the context, old or new lines of the diff")
	cmd.Flags().BoolS("z", "z", false, "do not munge pathnames and use NULs as output field terminators")
	rootCmd.AddCommand(cmd)

	cmd.Flag("color-moved").NoOptDefVal = "default"
	cmd.Flag("color-moved-ws").NoOptDefVal = " "
	cmd.Flag("word-diff").NoOptDefVal = "plain"

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"color":          carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"color-moved":    ActionColorMovedModes(),
		"color-moved-ws": ActionColorMovedWsModes(),
		"diff-algorithm": ActionDiffAlgorithms(),
		"output":         carapace.ActionFiles(),
		"submodule":      carapace.ActionValues("short", "long", "log"),
		"word-diff":      ActionWordDiffModes(),
		"ws-error-highlight": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionWsErrorHighlightModes().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(cmd).DashAnyCompletion(
		carapace.ActionFiles(),
	)

}

func ActionDiffAlgorithms() carapace.Action {
	return carapace.ActionValuesDescribed(
		"myers", "The basic greedy diff algorithm",
		"minimal", "Spend extra time to make sure the smallest possible diff is produced",
		"patience",
		"Use patience diff algorithm when generating patches",
		"histogram", "This algorithm extends the patience algorithm to support low-occurrence common elements",
	)
}

func ActionColorMovedModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"no", "Moved lines are not highlighted",
		"default", "default mode",
		"plain", "plain mode",
		"blocks", "greedily detects blocks",
		"zebra", "Blocks of moved text are detected as in blocks mode",
		"dimmed-zebra", "Similar to zebra, but additional dimming of uninteresting parts of moved code",
	)
}

func ActionColorMovedWsModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"no", "Do not ignore whitespace when performing move detection",
		"ignore-space-at-eol", "Ignore changes in whitespace at EOL",
		"ignore-space-change", "Ignore changes in amount of whitespace.",
		"ignore-all-space", "Ignore whitespace when comparing lines.",
		"allow-indentation-change", "Initially ignore any whitespace in the move detection",
	)
}

func ActionWordDiffModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"color", "Highlight changed words using only colors",
		"plain", "Show words as [-removed-] and {+added+}",
		"porcelain", "Use a special line-based format intended for script consumption",
		"none", "Disable word diff again",
	)
}

func ActionWsErrorHighlightModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"context", "context lines",
		"old", "old lines",
		"new", "new linex",
		"none", "reset previous values",
		"default", "reset to new",
		"all", "shorthand for old,new,context",
	)
}
