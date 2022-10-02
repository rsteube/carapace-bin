package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Updates configuration with the value of a given key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()
	config_setCmd.Flags().BoolP("global", "g", false, "write to global ~/.config/glab-cli/config.yml file rather than the repository .glab-cli/config/config")
	config_setCmd.Flags().StringP("host", "h", "", "Set per-host setting")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionConfigHosts(),
	})

	carapace.Gen(config_setCmd).PositionalCompletion(
		action.ActionConfigKeys(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionConfigValues(c.Args[0])
		}),
	)
}
