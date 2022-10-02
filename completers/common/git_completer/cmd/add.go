package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file contents to the index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolP("all", "A", false, "Update the index not only where the working tree has a file matching <pathspec> but also where the index already has an entry.")
	addCmd.Flags().String("chmod", "", "Override the executable bit of the added files.")
	addCmd.Flags().BoolP("dry-run", "n", false, "Don't actually add the file(s), just show if they exist and/or will be ignored.")
	addCmd.Flags().BoolP("edit", "e", false, "Open the diff vs. the index in an editor and let the user edit it.")
	addCmd.Flags().BoolP("force", "f", false, "Allow adding otherwise ignored files.")
	addCmd.Flags().Bool("ignore-errors", false, "If some files could not be added because of errors indexing them, do not abort the operation, but continue adding the others.")
	addCmd.Flags().Bool("ignore-missing", false, "This option can only be used together with --dry-run.")
	addCmd.Flags().Bool("ignore-removal", false, "Update the index by adding new files that are unknown to the index and files modified in the working tree, but ignore files that have been removed from the working tree.")
	addCmd.Flags().BoolP("intent-to-add", "N", false, "Record only the fact that the path will be added later. ")
	addCmd.Flags().BoolP("interactive", "i", false, "Add modified contents in the working tree interactively to the index.")
	addCmd.Flags().Bool("no-all", false, "Update the index by adding new files that are unknown to the index and files modified in the working tree, but ignore files that have been removed from the working tree.")
	addCmd.Flags().Bool("no-ignore-removal", false, "Update the index not only where the working tree has a file matching <pathspec> but also where the index already has an entry.")
	addCmd.Flags().Bool("no-warn-embedded-repo", false, "By default, git add will warn when adding an embedded repository to the index without using git submodule add to create an entry in .gitmodules.")
	addCmd.Flags().BoolP("patch", "p", false, "Interactively choose hunks of patch between the index and the work tree and add them to the index.")
	addCmd.Flags().Bool("pathspec-file-nul", false, "Pathspec elements are separated with NUL character and all other characters are taken literally.")
	addCmd.Flags().String("pathspec-from-file", "", "Pathspec is passed in <file> instead of commandline args.")
	addCmd.Flags().Bool("refresh", false, "Don't add the file(s), but only refresh their stat() information in the index.")
	addCmd.Flags().Bool("renormalize", false, "Apply the \"clean\" process freshly to all tracked files to forcibly add them again to the index.")
	addCmd.Flags().BoolP("update", "u", false, "Update the index just where it already has an entry matching <pathspec>.")
	addCmd.Flags().BoolP("verbose", "v", false, "Be verbose.")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"pathspec-from-file": carapace.ActionFiles(),
	})

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionChanges(git.ChangeOpts{Unstaged: true}).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
