package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_renameCmd).Standalone()

	remoteCmd.AddCommand(remote_renameCmd)

	carapace.Gen(remote_renameCmd).PositionalCompletion(
		git.ActionRemotes(),
		git.ActionRemotes(),
	)
}
