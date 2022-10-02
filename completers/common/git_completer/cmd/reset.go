package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset current HEAD to the specified state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetCmd).Standalone()

	resetCmd.Flags().Bool("hard", false, "reset HEAD, index and working tree")
	resetCmd.Flags().BoolP("intent-to-add", "N", false, "record only the fact that removed paths will be added later")
	resetCmd.Flags().Bool("keep", false, "reset HEAD but keep local changes")
	resetCmd.Flags().Bool("merge", false, "reset HEAD, index and working tree")
	resetCmd.Flags().Bool("mixed", false, "reset HEAD and index")
	resetCmd.Flags().BoolP("patch", "p", false, "select hunks interactively")
	resetCmd.Flags().Bool("pathspec-file-nul", false, "pathspec elements are separated with NUL character")
	resetCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	resetCmd.Flags().BoolP("quiet", "q", false, "be quiet, only report errors")
	resetCmd.Flags().String("recurse-submodules", "", "control recursive updating of submodules")
	resetCmd.Flags().Bool("soft", false, "reset only HEAD")
	rootCmd.AddCommand(resetCmd)

	resetCmd.Flag("recurse-submodules").NoOptDefVal = " "

	carapace.Gen(resetCmd).FlagCompletion(carapace.ActionMap{
		"pathspec-from-file": carapace.ActionFiles(),
	})

	carapace.Gen(resetCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(resetCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if resetCmd.Flags().ArgsLenAtDash() != -1 {
				return carapace.ActionFiles()
			}
			return carapace.ActionValues()
		}),
	)
}
