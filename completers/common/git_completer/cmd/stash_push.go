package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stash_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "save your local modifications to a new stash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_pushCmd).Standalone()

	stash_pushCmd.Flags().BoolP("all", "a", false, "also stash ignored and untracked")
	stash_pushCmd.Flags().BoolP("include-untracked", "u", false, "also stash untracked")
	stash_pushCmd.Flags().BoolP("keep-index", "k", false, "keep changed added to index")
	stash_pushCmd.Flags().StringP("message", "m", "", "set description")
	stash_pushCmd.Flags().Bool("no-keep-index", false, "also apply to index")
	stash_pushCmd.Flags().BoolP("patch", "p", false, "interactively select hunks between HEAD and working tree")
	stash_pushCmd.Flags().Bool("pathspec-file-nul", false, "pathspec elemts are seperated by NUL")
	stash_pushCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	stash_pushCmd.Flags().BoolP("quiet", "q", false, "suppress feedback messages")
	stashCmd.AddCommand(stash_pushCmd)

	carapace.Gen(stash_pushCmd).FlagCompletion(carapace.ActionMap{
		"pathspec-from-file": carapace.ActionFiles(),
	})
}
