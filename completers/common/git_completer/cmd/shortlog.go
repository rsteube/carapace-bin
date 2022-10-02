package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var shortlogCmd = &cobra.Command{
	Use:   "shortlog",
	Short: "Summarize 'git log' output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shortlogCmd).Standalone()

	shortlogCmd.Flags().String("after", "", "Show commits more recent than a specific date.")
	shortlogCmd.Flags().Bool("all", false, "Pretend as if all the refs in refs/, along with HEAD, are listed on the command line as <commit>.")
	shortlogCmd.Flags().Bool("all-match", false, "Limit the commits output to ones that match all given --grep, instead of ones that match at least one.")
	shortlogCmd.Flags().Bool("alternate-refs", false, "Pretend as if all objects mentioned as ref tips of alternate repositories were listed on the command line.")
	shortlogCmd.Flags().Bool("ancestry-path", false, "Limit the displayed commits to those directly on the ancestry chain between the “from” and “to” commits in the given commit range.")
	shortlogCmd.Flags().String("author", "", "Limit the commits output to ones with author header lines that match the specified pattern.")
	shortlogCmd.Flags().Bool("basic-regexp", false, "Consider the limiting patterns to be basic regular expressions; this is the default.")
	shortlogCmd.Flags().String("before", "", "Show commits older than a specific date.")
	shortlogCmd.Flags().Bool("bisect", false, "Pretend as if the bad bisection ref refs/bisect/bad was listed and as if it was followed by --not and the good bisection refs refs/bisect/good-* on the command line.")
	shortlogCmd.Flags().Bool("boundary", false, "Output excluded boundary commits. Boundary commits are prefixed with -.")
	shortlogCmd.Flags().Bool("cherry", false, "A synonym for --right-only --cherry-mark --no-merges.")
	shortlogCmd.Flags().Bool("cherry-mark", false, "Like --cherry-pick (see below) but mark equivalent commits with = rather than omitting them, and inequivalent ones with +.")
	shortlogCmd.Flags().Bool("cherry-pick", false, "Omit any commit that introduces the same change as another commit on the “other side” when the set of commits are limited with symmetric difference.")
	shortlogCmd.Flags().BoolP("committer", "c", false, "This is an alias for --group=committer.")
	shortlogCmd.Flags().Bool("dense", false, "Only the selected commits are shown, plus some to have a meaningful history.")
	shortlogCmd.Flags().BoolP("email", "e", false, "Show the email address of each author.")
	shortlogCmd.Flags().String("exclude", "", "Do not include refs matching <glob-pattern> that the next --all, --branches, --tags, --remotes, or --glob would otherwise consider.")
	shortlogCmd.Flags().BoolP("extended-regexp", "E", false, "Consider the limiting patterns to be extended regular expressions instead of the default basic regular expressions.")
	shortlogCmd.Flags().Bool("first-parent", false, "Follow only the first parent commit upon seeing a merge commit.")
	shortlogCmd.Flags().BoolP("fixed-strings", "F", false, "Consider the limiting patterns to be fixed strings.")
	shortlogCmd.Flags().String("format", "", "Instead of the commit subject, use some other information to describe each commit.")
	shortlogCmd.Flags().Bool("full-history", false, "Same as the default mode, but does not prune some history.")
	shortlogCmd.Flags().String("glob", "", "Pretend as if all the refs matching shell glob <glob-pattern> are listed on the command line as <commit>.")
	shortlogCmd.Flags().String("grep", "", "Limit the commits output to ones with log message that matches the specified pattern.")
	shortlogCmd.Flags().String("grep-reflog", "", "Limit the commits output to ones with reflog entries that match the specified pattern.")
	shortlogCmd.Flags().String("group", "", "Group commits based on <type>.")
	shortlogCmd.Flags().BoolP("help", "h", false, "Show help.")
	shortlogCmd.Flags().Bool("ignore-missing", false, "Upon seeing an invalid object name in the input, pretend as if the bad input was not given.")
	shortlogCmd.Flags().Bool("invert-grep", false, "Limit the commits output to ones with log message that do not match the pattern specified with --grep=<pattern>.")
	shortlogCmd.Flags().Bool("left-only", false, "List only commits on the respective side of a symmetric difference.")
	shortlogCmd.Flags().String("max-count", "", "Limit the number of commits to output.")
	shortlogCmd.Flags().String("max-parents", "", "Show only commits which have at most that many parent commits.")
	shortlogCmd.Flags().Bool("merge", false, "After a failed merge, show refs that touch files having a conflict and don’t exist on all heads to merge.")
	shortlogCmd.Flags().Bool("merges", false, "Print only merge commits.")
	shortlogCmd.Flags().String("min-parents", "", "Show only commits which have at least that many parent commits.")
	shortlogCmd.Flags().Bool("no-max-parents", false, "")
	shortlogCmd.Flags().Bool("no-merges", false, "Do not print commits with more than one parent.")
	shortlogCmd.Flags().Bool("no-min-parents", false, "")
	shortlogCmd.Flags().Bool("not", false, "Reverses the meaning of the ^ prefix (or lack thereof) for all following revision specifiers, up to the next --not.")
	shortlogCmd.Flags().BoolP("numbered", "n", false, "Sort output according to the number of commits per author instead of author alphabetic order.")
	shortlogCmd.Flags().BoolP("perl-regexp", "P", false, "Consider the limiting patterns to be Perl-compatible regular expressions.")
	shortlogCmd.Flags().Bool("reflog", false, "Pretend as if all objects mentioned by reflogs are listed on the command line as <commit>.")
	shortlogCmd.Flags().BoolP("regexp-ignore-case", "i", false, "Match the regular expression limiting patterns without regard to letter case.")
	shortlogCmd.Flags().String("remotes", "", "Pretend as if all the refs in refs/remotes are listed on the command line as <commit>.")
	shortlogCmd.Flags().Bool("remove-empty", false, "Stop when a given path disappears from the tree.")
	shortlogCmd.Flags().Bool("right-only", false, "List only commits on the respective side of a symmetric difference.")
	shortlogCmd.Flags().Bool("show-pulls", false, "Include all commits from the default mode, but also any merge commits that are not TREESAME to the first parent but are TREESAME to a later parent.")
	shortlogCmd.Flags().Bool("simplify-by-decoration", false, "Commits that are referred by some branch or tag are selected.")
	shortlogCmd.Flags().Bool("simplify-merges", false, "First, build a history graph in the same way that --full-history with parent rewriting does.")
	shortlogCmd.Flags().String("since", "", "Show commits more recent than a specific date.")
	shortlogCmd.Flags().Bool("single-worktree", false, "By default, all working trees will be examined by the following options when there are more than one.")
	shortlogCmd.Flags().Bool("sparse", false, "All commits in the simplified history are shown.")
	shortlogCmd.Flags().Bool("stdin", false, "In addition to the <commit> listed on the command line, read them from the standard input.")
	shortlogCmd.Flags().BoolP("summary", "s", false, "Suppress commit description and provide a commit count summary only.")
	shortlogCmd.Flags().String("tags", "", "Pretend as if all the refs in refs/tags are listed on the command line as <commit>.")
	shortlogCmd.Flags().String("until", "", "Show commits older than a specific date.")
	shortlogCmd.Flags().StringS("w", "w", "", "Linewrap the output by wrapping each line at width.")
	shortlogCmd.Flags().BoolP("walk-reflogs", "g", false, "Instead of walking the commit ancestry chain, walk reflog entries from the most recent one to older ones.")
	rootCmd.AddCommand(shortlogCmd)

	shortlogCmd.Flag("glob").NoOptDefVal = " "
	shortlogCmd.Flag("remotes").NoOptDefVal = " "
	shortlogCmd.Flag("tags").NoOptDefVal = " "

	carapace.Gen(shortlogCmd).FlagCompletion(carapace.ActionMap{
		"after":  time.ActionDate(),
		"author": git.ActionAuthors(),
		"before": time.ActionDate(),
		"group": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"author", "commits are grouped by author",
					"committer", "commits are grouped by committer",
					"trailer:", "a case-insensitive commit message trailer",
				)
			default:
				return carapace.ActionValues()
			}
		}),
		"since": time.ActionDate(),
		"until": time.ActionDate(),
	})

	carapace.Gen(shortlogCmd).PositionalAnyCompletion(
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
