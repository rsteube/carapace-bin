package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Join two or more development histories together",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	mergeCmd.Flags().Bool("abort", false, "abort the current in-progress merge")
	mergeCmd.Flags().Bool("allow-unrelated-histories", false, "allow merging unrelated histories")
	mergeCmd.Flags().Bool("autostash", false, "automatically stash/stash pop before and after")
	mergeCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	mergeCmd.Flags().Bool("commit", false, "perform a commit if the merge succeeds (default)")
	mergeCmd.Flags().Bool("continue", false, "continue the current in-progress merge")
	mergeCmd.Flags().BoolP("edit", "e", false, "edit message before committing")
	mergeCmd.Flags().Bool("ff", false, "allow fast-forward (default)")
	mergeCmd.Flags().Bool("ff-only", false, "abort if fast-forward is not possible")
	mergeCmd.Flags().StringP("file", "F", "", "read message from file")
	mergeCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	mergeCmd.Flags().String("log", "", "add (at most <n>) entries from shortlog to merge commit message")
	mergeCmd.Flags().StringP("message", "m", "", "merge commit message")
	mergeCmd.Flags().BoolS("n", "n", false, "do not show a diffstat at the end of the merge")
	mergeCmd.Flags().Bool("no-verify", false, "bypass pre-merge-commit and commit-msg hooks")
	mergeCmd.Flags().Bool("overwrite-ignore", false, "update ignored files (default)")
	mergeCmd.Flags().Bool("progress", false, "force progress reporting")
	mergeCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	mergeCmd.Flags().Bool("quit", false, "--abort but leave index and working tree alone")
	mergeCmd.Flags().Bool("rerere-autoupdate", false, "update the index with reused conflict resolution if possible")
	mergeCmd.Flags().Bool("signoff", false, "add Signed-off-by:")
	mergeCmd.Flags().Bool("squash", false, "create a single commit instead of doing a merge")
	mergeCmd.Flags().Bool("stat", false, "show a diffstat at the end of the merge")
	mergeCmd.Flags().StringP("strategy", "s", "", "merge strategy to use")
	mergeCmd.Flags().StringP("strategy-option", "X", "", "option for selected merge strategy")
	mergeCmd.Flags().Bool("summary", false, "(synonym to --stat)")
	mergeCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	mergeCmd.Flags().Bool("verify-signatures", false, "verify that the named commit has a valid GPG signature")
	rootCmd.AddCommand(mergeCmd)

	mergeCmd.Flag("gpg-sign").NoOptDefVal = " "

	carapace.Gen(mergeCmd).FlagCompletion(carapace.ActionMap{
		"cleanup":  git.ActionCleanupMode(),
		"file":     carapace.ActionFiles(),
		"gpg-sign": os.ActionGpgKeyIds(),
		"strategy": git.ActionMergeStrategy(),
		"strategy-option": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionMergeStrategyOptions(mergeCmd.Flag("strategy").Value.String())
		}),
	})

	carapace.Gen(mergeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRefs(git.RefOption{}.Default()).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
