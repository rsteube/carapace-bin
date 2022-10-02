package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var isEnabledCmd = &cobra.Command{
	Use:   "is-enabled",
	Short: "Check whether unit files are enabled",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isEnabledCmd).Standalone()

	rootCmd.AddCommand(isEnabledCmd)

	carapace.Gen(isEnabledCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
