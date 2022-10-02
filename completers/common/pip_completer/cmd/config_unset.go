package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pip"
	"github.com/spf13/cobra"
)

var config_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset the value associated with name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_unsetCmd).Standalone()
	configCmd.AddCommand(config_unsetCmd)

	carapace.Gen(config_unsetCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return pip.ActionConfigValues(pip.ConfigOpts{
				Global: configCmd.Flag("global").Changed,
				Site:   configCmd.Flag("site").Changed,
				User:   configCmd.Flag("user").Changed,
			})
		}),
	)
}
