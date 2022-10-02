package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_setBranchesCmd = &cobra.Command{
	Use:   "set-branches",
	Short: "Changes the list of branches tracked",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_setBranchesCmd).Standalone()

	remote_setBranchesCmd.Flags().Bool("add", false, "add branch")
	remoteCmd.AddCommand(remote_setBranchesCmd)

	carapace.Gen(remote_setBranchesCmd).PositionalCompletion(
		git.ActionRemotes(),
	)

	carapace.Gen(remote_setBranchesCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRemoteBranches(c.Args[0])
		}),
	)
}
