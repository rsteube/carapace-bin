package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable one or more unit files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableCmd).Standalone()

	rootCmd.AddCommand(enableCmd)

	carapace.Gen(enableCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
			}),
			carapace.ActionFiles(),
		).ToA(),
	)
}
