package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_setUrlCmd = &cobra.Command{
	Use:   "set-url",
	Short: "Changes URLs for the remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_setUrlCmd).Standalone()

	remote_setUrlCmd.Flags().Bool("add", false, "add URL")
	remote_setUrlCmd.Flags().Bool("delete", false, "delete URLs")
	remote_setUrlCmd.Flags().Bool("push", false, "manipulate push URLs")
	remoteCmd.AddCommand(remote_setUrlCmd)

	carapace.Gen(remote_setUrlCmd).PositionalCompletion(
		git.ActionRemotes(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				git.ActionRemoteUrls(c.Args[0]),
				git.ActionRepositorySearch(),
			).ToA()
		}),
	)
}
