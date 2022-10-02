package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Record changes to the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().Bool("ahead-behind", false, "compute full ahead/behind values")
	commitCmd.Flags().BoolP("all", "a", false, "commit all changed files")
	commitCmd.Flags().Bool("amend", false, "amend previous commit")
	commitCmd.Flags().String("author", "", "override author for commit")
	commitCmd.Flags().Bool("branch", false, "show branch information")
	commitCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	commitCmd.Flags().String("date", "", "override date for commit")
	commitCmd.Flags().Bool("dry-run", false, "show what would be committed")
	commitCmd.Flags().BoolP("edit", "e", false, "force edit of commit")
	commitCmd.Flags().StringP("file", "F", "", "read message from file")
	commitCmd.Flags().String("fixup", "", "use autosquash formatted message to fixup specified commit")
	commitCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	commitCmd.Flags().BoolP("include", "i", false, "add specified files to index for commit")
	commitCmd.Flags().Bool("interactive", false, "interactively add files")
	commitCmd.Flags().Bool("long", false, "show status in long format (default)")
	commitCmd.Flags().StringP("message", "m", "", "commit message")
	commitCmd.Flags().Bool("no-post-rewrite", false, "bypass post-rewrite hook")
	commitCmd.Flags().BoolP("no-verify", "n", false, "bypass pre-commit and commit-msg hooks")
	commitCmd.Flags().BoolP("null", "z", false, "terminate entries with NUL")
	commitCmd.Flags().BoolP("only", "o", false, "commit only specified files")
	commitCmd.Flags().BoolP("patch", "p", false, "interactively add changes")
	commitCmd.Flags().Bool("pathspec-file-nul", false, "pathspec elements are separated with NUL character")
	commitCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	commitCmd.Flags().Bool("porcelain", false, "machine-readable output")
	commitCmd.Flags().BoolP("quiet", "q", false, "suppress summary after successful commit")
	commitCmd.Flags().StringP("reedit-message", "c", "", "reuse and edit message from specified commit")
	commitCmd.Flags().Bool("reset-author", false, "the commit is authored by me now (used with -C/-c/--amend)")
	commitCmd.Flags().StringP("reuse-message", "C", "", "reuse message from specified commit")
	commitCmd.Flags().Bool("short", false, "show status concisely")
	commitCmd.Flags().BoolP("signoff", "s", false, "add Signed-off-by:")
	commitCmd.Flags().String("squash", "", "use autosquash formatted message to squash specified commit")
	commitCmd.Flags().Bool("status", false, "include status in commit message template")
	commitCmd.Flags().StringP("template", "t", "", "use specified template file")
	commitCmd.Flags().StringP("untracked-files", "u", "", "show untracked files")
	commitCmd.Flags().BoolP("verbose", "v", false, "show diff in commit message template")
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flag("gpg-sign").NoOptDefVal = " "
	commitCmd.Flag("untracked-files").NoOptDefVal = " "

	carapace.Gen(commitCmd).FlagCompletion(carapace.ActionMap{
		"cleanup":            git.ActionCleanupMode(),
		"file":               carapace.ActionFiles(),
		"fixup":              git.ActionRefs(git.RefOption{Commits: 100}),
		"gpg-sign":           os.ActionGpgKeyIds(),
		"pathspec-from-file": carapace.ActionFiles(),
		"reedit-message":     git.ActionRefs(git.RefOption{Commits: 100}),
		"reuse-message":      git.ActionRefs(git.RefOption{Commits: 100}),
		"squash":             git.ActionRefs(git.RefOption{Commits: 100}),
		"template":           carapace.ActionFiles(),
		"untracked-files":    carapace.ActionValues("all", "normal", "no"),
	})
}
