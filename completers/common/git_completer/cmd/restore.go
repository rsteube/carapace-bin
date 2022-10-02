package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore working tree files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restoreCmd).Standalone()

	restoreCmd.Flags().String("conflict", "", "conflict style (merge or diff3)")
	restoreCmd.Flags().Bool("ignore-skip-worktree-bits", false, "do not limit pathspecs to sparse entries only")
	restoreCmd.Flags().Bool("ignore-unmerged", false, "ignore unmerged entries")
	restoreCmd.Flags().BoolP("merge", "m", false, "perform a 3-way merge with the new branch")
	restoreCmd.Flags().BoolP("ours", "2", false, "checkout our version for unmerged files")
	restoreCmd.Flags().Bool("overlay", false, "use overlay mode")
	restoreCmd.Flags().BoolP("patch", "p", false, "select hunks interactively")
	restoreCmd.Flags().Bool("pathspec-file-nul", false, "pathspec elements are separated with NUL character")
	restoreCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	restoreCmd.Flags().Bool("progress", false, "force progress reporting")
	restoreCmd.Flags().BoolP("quiet", "q", false, "suppress progress reporting")
	restoreCmd.Flags().String("recurse-submodules", "", "control recursive updating of submodules")
	restoreCmd.Flags().StringP("source", "s", "", "which tree-ish to checkout from")
	restoreCmd.Flags().BoolP("staged", "S", false, "restore the index")
	restoreCmd.Flags().BoolP("theirs", "3", false, "checkout their version for unmerged files")
	restoreCmd.Flags().BoolP("worktree", "W", false, "restore the working tree (default)")
	rootCmd.AddCommand(restoreCmd)

	carapace.Gen(restoreCmd).FlagCompletion(carapace.ActionMap{
		"conflict":           carapace.ActionValues("merge", "diff3"),
		"pathspec-from-file": carapace.ActionFiles(),
		"recurse-submodules": git.ActionRefs(git.RefOption{}.Default()),
		"source":             git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(restoreCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if restoreCmd.Flag("staged").Changed {
				return git.ActionChanges(git.ChangeOpts{Staged: true}).Invoke(c).Filter(c.Args).ToA()
			}
			return git.ActionChanges(git.ChangeOpts{Unstaged: true}).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
