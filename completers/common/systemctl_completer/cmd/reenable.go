package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reenableCmd = &cobra.Command{
	Use:   "reenable",
	Short: "Reenable one or more unit files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reenableCmd).Standalone()

	rootCmd.AddCommand(reenableCmd)

	carapace.Gen(reenableCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
