package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/circleci_completer/cmd/action"
	"github.com/spf13/cobra"
)

var orb_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show the meta-data of an orb",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_infoCmd).Standalone()
	orbCmd.AddCommand(orb_infoCmd)

	carapace.Gen(orb_infoCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionOrbs().Invoke(c).ToMultiPartsA("/")
		}),
	)
}
