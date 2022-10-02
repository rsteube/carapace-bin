package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listDependenciesCmd = &cobra.Command{
	Use:   "list-dependencies",
	Short: "Recursively show units which are required or wanted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listDependenciesCmd).Standalone()

	rootCmd.AddCommand(listDependenciesCmd)

	carapace.Gen(listDependenciesCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
