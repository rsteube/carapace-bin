package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert some existing commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revertCmd).Standalone()

	revertCmd.Flags().Bool("abort", false, "cancel revert or cherry-pick sequence")
	revertCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	revertCmd.Flags().Bool("continue", false, "resume revert or cherry-pick sequence")
	revertCmd.Flags().BoolP("edit", "e", false, "edit the commit message")
	revertCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	revertCmd.Flags().StringP("mainline", "m", "", "select mainline parent")
	revertCmd.Flags().BoolP("no-commit", "n", false, "don't automatically commit")
	revertCmd.Flags().Bool("quit", false, "end revert or cherry-pick sequence")
	revertCmd.Flags().Bool("rerere-autoupdate", false, "update the index with reused conflict resolution if possible")
	revertCmd.Flags().BoolP("signoff", "s", false, "add Signed-off-by:")
	revertCmd.Flags().Bool("skip", false, "skip current commit and continue")
	revertCmd.Flags().String("strategy", "", "merge strategy")
	revertCmd.Flags().StringP("strategy-option", "X", "", "option for merge strategy")
	rootCmd.AddCommand(revertCmd)

	revertCmd.Flag("gpg-sign").NoOptDefVal = " "

	carapace.Gen(revertCmd).FlagCompletion(carapace.ActionMap{
		"cleanup":  git.ActionCleanupMode(),
		"gpg-sign": os.ActionGpgKeyIds(),
		"strategy": git.ActionMergeStrategy(),
		"strategy-option": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionMergeStrategyOptions(revertCmd.Flag("strategy").Value.String())
		}),
	})

	carapace.Gen(revertCmd).PositionalAnyCompletion(git.ActionRefs(git.RefOption{}.Default()))
}
