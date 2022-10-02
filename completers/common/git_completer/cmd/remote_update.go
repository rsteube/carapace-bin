package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Fetch updates for remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_updateCmd).Standalone()

	remote_updateCmd.Flags().BoolP("prune", "p", false, "prune remotes after fetching")
	remoteCmd.AddCommand(remote_updateCmd)

	carapace.Gen(remote_updateCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRemotes().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
