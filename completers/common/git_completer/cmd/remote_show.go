package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Give some inforemation about the remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_showCmd).Standalone()

	remote_showCmd.Flags().BoolS("n", "n", false, "do not query remotes")
	remoteCmd.AddCommand(remote_showCmd)

	carapace.Gen(remote_showCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
