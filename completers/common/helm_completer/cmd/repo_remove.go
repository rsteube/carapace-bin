package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove one or more chart repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	repoCmd.AddCommand(repo_removeCmd)

	carapace.Gen(repo_removeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionRepositories().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
