package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Print a list of configuration keys and values",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()
	config_listCmd.Flags().StringP("host", "h", "", "Get per-host configuration")
	configCmd.AddCommand(config_listCmd)

	carapace.Gen(config_listCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionConfigHosts(),
	})
}
