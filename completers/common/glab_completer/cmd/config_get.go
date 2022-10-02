package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Prints the value of a given configuration key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()
	config_getCmd.Flags().BoolP("global", "g", false, "Read from global config file (~/.config/glab-cli/config.yml). [Default: looks through Environment variables → Local → Global]")
	config_getCmd.Flags().StringP("host", "h", "", "Get per-host setting")
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionConfigHosts(),
	})

	carapace.Gen(config_getCmd).PositionalCompletion(
		action.ActionConfigKeys(),
	)
}
