package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var addRequiresCmd = &cobra.Command{
	Use:   "add-requires",
	Short: "Add 'Requires' dependency for the target on specified one or more units",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addRequiresCmd).Standalone()

	rootCmd.AddCommand(addRequiresCmd)

	carapace.Gen(addRequiresCmd).PositionalCompletion(
		action.ActionTargets(),
	)

	carapace.Gen(addRequiresCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
