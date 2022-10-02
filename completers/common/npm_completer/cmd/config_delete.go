package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes the specified keys from all configuration files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteCmd).Standalone()

	configCmd.AddCommand(config_deleteCmd)

	carapace.Gen(config_deleteCmd).PositionalAnyCompletion(
		action.ActionConfigKeys(config_deleteCmd),
	)
}
