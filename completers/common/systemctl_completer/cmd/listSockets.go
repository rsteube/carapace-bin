package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listSocketsCmd = &cobra.Command{
	Use:   "list-sockets",
	Short: "List socket units currently in memory, ordered by address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listSocketsCmd).Standalone()

	rootCmd.AddCommand(listSocketsCmd)

	carapace.Gen(listSocketsCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
