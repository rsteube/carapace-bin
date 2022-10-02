package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var gist_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a gist locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gist_cloneCmd).Standalone()
	gistCmd.AddCommand(gist_cloneCmd)

	carapace.Gen(gist_cloneCmd).PositionalCompletion(
		action.ActionGists(gist_cloneCmd),
		carapace.ActionDirectories(),
	)
}
