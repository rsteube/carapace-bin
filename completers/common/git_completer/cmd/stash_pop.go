package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_popCmd = &cobra.Command{
	Use:     "pop",
	Aliases: []string{"apply"},
	Short:   "remove a single stashed state",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_popCmd).Standalone()
	stash_popCmd.Flags().Bool("index", false, "try to reinstate index changes as well")
	stash_popCmd.Flags().BoolP("quiet", "q", false, "suppress feedback messages")

	stashCmd.AddCommand(stash_popCmd)

	carapace.Gen(stash_popCmd).PositionalCompletion(git.ActionStashes())
}
