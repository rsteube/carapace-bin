package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_getUrlCmd = &cobra.Command{
	Use:   "get-url",
	Short: "Retrieves the URLs for a remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_getUrlCmd).Standalone()

	remote_getUrlCmd.Flags().Bool("all", false, "return all URLs")
	remote_getUrlCmd.Flags().Bool("push", false, "query push URLs rather than fetch URLs")
	remoteCmd.AddCommand(remote_getUrlCmd)

	carapace.Gen(remote_getUrlCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
