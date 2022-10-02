package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "remove a single stash entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_dropCmd).Standalone()
	stash_dropCmd.Flags().BoolP("quiet", "q", false, "suppress feedback messages")

	stashCmd.AddCommand(stash_dropCmd)

	carapace.Gen(stash_dropCmd).PositionalCompletion(git.ActionStashes())
}
