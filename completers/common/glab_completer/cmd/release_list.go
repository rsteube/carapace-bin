package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List releases in a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_listCmd).Standalone()
	release_listCmd.Flags().StringP("tag", "t", "", "Filter releases by tag <name>")
	releaseCmd.AddCommand(release_listCmd)

	carapace.Gen(release_listCmd).FlagCompletion(carapace.ActionMap{
		"tag": action.ActionTags(release_listCmd),
	})
}
