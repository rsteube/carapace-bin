package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var addWantsCmd = &cobra.Command{
	Use:   "add-wants",
	Short: "Add 'Wants' dependency for the target on specified one or more units",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addWantsCmd).Standalone()

	rootCmd.AddCommand(addWantsCmd)

	carapace.Gen(addWantsCmd).PositionalCompletion(
		action.ActionTargets(),
	)

	carapace.Gen(addWantsCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
