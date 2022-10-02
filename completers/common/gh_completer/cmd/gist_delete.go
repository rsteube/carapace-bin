package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var gist_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a gist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gist_deleteCmd).Standalone()
	gistCmd.AddCommand(gist_deleteCmd)

	carapace.Gen(gist_deleteCmd).PositionalCompletion(
		action.ActionGists(gist_deleteCmd),
	)
}
