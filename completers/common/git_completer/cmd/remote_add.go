package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_addCmd).Standalone()

	remote_addCmd.Flags().BoolP("fetch", "f", false, "fetch the remote branches")
	remote_addCmd.Flags().StringP("master", "m", "", "master branch")
	remote_addCmd.Flags().String("mirror", "", "set up remote as a mirror to push to or fetch from")
	remote_addCmd.Flags().Bool("tags", false, "import all tags and associated objects when fetching")
	remote_addCmd.Flags().StringArrayP("track", "t", []string{}, "branch(es) to track")

	remote_addCmd.Flag("mirror").NoOptDefVal = " "

	remoteCmd.AddCommand(remote_addCmd)

	carapace.Gen(remote_addCmd).FlagCompletion(carapace.ActionMap{
		"master": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 1 {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[1], Branches: true})
			} else {
				return carapace.ActionValues()
			}
		}),
		"mirror": carapace.ActionValues("push", "fetch"),
		"track": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 1 {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: c.Args[1], Branches: true})
			} else {
				return carapace.ActionValues()
			}
		}),
	})

	carapace.Gen(remote_addCmd).PositionalCompletion(
		git.ActionRemotes(),
		git.ActionRepositorySearch(),
	)
}
