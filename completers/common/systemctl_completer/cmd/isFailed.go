package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var isFailedCmd = &cobra.Command{
	Use:   "is-failed",
	Short: "Check whether units are failed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isFailedCmd).Standalone()

	rootCmd.AddCommand(isFailedCmd)

	carapace.Gen(isFailedCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
