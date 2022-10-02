package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "create and check out a new branch with the stashes changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_branchCmd).Standalone()

	stashCmd.AddCommand(stash_branchCmd)

	carapace.Gen(stash_branchCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{LocalBranches: true}),
		git.ActionStashes(),
	)
}
