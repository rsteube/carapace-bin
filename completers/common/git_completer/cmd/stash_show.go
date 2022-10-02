package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the changes recorded in the stash entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_showCmd).Standalone()
	stash_showCmd.Flags().Bool("include-untracked", false, "include untracked")
	stash_showCmd.Flags().Bool("only-untracked", false, "only untracked")
	addDiffFlags(stash_showCmd)

	stashCmd.AddCommand(stash_showCmd)

	carapace.Gen(stash_showCmd).PositionalCompletion(git.ActionStashes())
}
