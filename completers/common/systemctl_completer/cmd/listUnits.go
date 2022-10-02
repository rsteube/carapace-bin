package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listUnitsCmd = &cobra.Command{
	Use:   "list-units",
	Short: "List units currently in memory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listUnitsCmd).Standalone()

	rootCmd.AddCommand(listUnitsCmd)

	carapace.Gen(listUnitsCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
