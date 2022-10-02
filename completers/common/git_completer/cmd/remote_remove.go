package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_removeCmd).Standalone()

	remoteCmd.AddCommand(remote_removeCmd)

	carapace.Gen(remote_removeCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
