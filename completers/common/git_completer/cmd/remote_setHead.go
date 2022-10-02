package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_setHeadCmd = &cobra.Command{
	Use:   "set-head",
	Short: "Sets or deletes the default branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_setHeadCmd).Standalone()

	remote_setHeadCmd.Flags().BoolP("auto", "a", false, "set refs/remotes/<name>/HEAD according to remote")
	remote_setHeadCmd.Flags().BoolP("delete", "d", false, "delete refs/remotes/<name>/HEAD")
	remoteCmd.AddCommand(remote_setHeadCmd)

	carapace.Gen(remote_setHeadCmd).PositionalCompletion(
		git.ActionRemotes(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRemoteBranches(c.Args[0])
		}),
	)
}
