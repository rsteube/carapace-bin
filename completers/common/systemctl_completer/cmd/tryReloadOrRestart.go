package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tryReloadOrRestartCmd = &cobra.Command{
	Use:   "try-reload-or-restart",
	Short: "If active, reload one or more units, if supported, otherwise restart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tryReloadOrRestartCmd).Standalone()

	rootCmd.AddCommand(tryReloadOrRestartCmd)

	carapace.Gen(tryReloadOrRestartCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
