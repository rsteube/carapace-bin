package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var cherryPickCmd = &cobra.Command{
	Use:   "cherry-pick",
	Short: "Apply the changes introduced by some existing commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cherryPickCmd).Standalone()

	cherryPickCmd.Flags().Bool("abort", false, "cancel the operation and return to the pre-sequence state")
	cherryPickCmd.Flags().Bool("allow-empty", false, "preserve empty commits")
	cherryPickCmd.Flags().Bool("allow-empty-message", false, "allow empty commit message")
	cherryPickCmd.Flags().String("cleanup", "", "set commit message cleanup mode")
	cherryPickCmd.Flags().Bool("continue", false, "ontinue the operation in progress")
	cherryPickCmd.Flags().BoolP("edit", "e", false, "edit the commit message prior to committing")
	cherryPickCmd.Flags().Bool("ff", false, "fast forward")
	cherryPickCmd.Flags().StringP("gpg-sign", "S", "", "GPG-sign commits")
	cherryPickCmd.Flags().Bool("keep-redundant-commits", false, "preserve redundant commits")
	cherryPickCmd.Flags().StringP("mainline", "m", "", "specify parent number to pick relative change")
	cherryPickCmd.Flags().BoolP("no-commit", "n", false, "apply changes without committing them")
	cherryPickCmd.Flags().Bool("no-gpg-sign", false, "do not GPG-sign commits")
	cherryPickCmd.Flags().Bool("no-rerere-autoupdate", false, "disallow rerere mechanism to update the index")
	cherryPickCmd.Flags().Bool("quit", false, "forget about the current operation in progress")
	cherryPickCmd.Flags().Bool("rerere-autoupdate", false, "allow rerere mechanism to update the index")
	cherryPickCmd.Flags().BoolP("signoff", "s", false, "add Signed-off-by line at the end of the commit message")
	cherryPickCmd.Flags().Bool("skip", false, "skip the current commit")
	cherryPickCmd.Flags().String("strategy", "", "use the given merge strategy")
	cherryPickCmd.Flags().StringP("strategy-option", "X", "", "pass the merge strategy-specific option through to the merge strategy")
	cherryPickCmd.Flags().BoolS("x", "x", false, "append cherry-pick source to original commit message")
	rootCmd.AddCommand(cherryPickCmd)

	cherryPickCmd.Flag("gpg-sign").NoOptDefVal = " "

	carapace.Gen(cherryPickCmd).FlagCompletion(carapace.ActionMap{
		"cleanup":  git.ActionCleanupMode(),
		"gpg-sign": os.ActionGpgKeyIds(),
		"mainline": carapace.ActionValues("1", "2", "3", "4", "5"),
		"strategy": git.ActionMergeStrategy(),
		"strategy-option": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionMergeStrategyOptions(cherryPickCmd.Flag("strategy").Value.String())
		}),
	})

	carapace.Gen(cherryPickCmd).PositionalAnyCompletion(
		// TODO `...` divider not yet working
		carapace.ActionMultiParts("...", func(c carapace.Context) carapace.Action {
			if len(c.Parts) < 2 {
				return git.ActionRefs(git.RefOption{}.Default())
			}
			return carapace.ActionValues()
		}),
	)
}
