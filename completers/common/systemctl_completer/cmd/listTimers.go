package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listTimersCmd = &cobra.Command{
	Use:   "list-timers",
	Short: "List timer units currently in memory, ordered by next elapse",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listTimersCmd).Standalone()

	rootCmd.AddCommand(listTimersCmd)

	carapace.Gen(listTimersCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
