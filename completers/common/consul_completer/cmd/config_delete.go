package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a centralized config entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteCmd).Standalone()
	config_deleteCmd.Flags().String("kind", "", "The kind of configuration to delete.")
	config_deleteCmd.Flags().String("name", "", "The name of configuration to delete.")

	configCmd.AddCommand(config_deleteCmd)

	carapace.Gen(config_deleteCmd).FlagCompletion(carapace.ActionMap{
		"kind": action.ActionConfigKinds(),
		"name": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := config_deleteCmd.Flag("kind"); flag.Changed {
				return action.ActionConfigEntries(flag.Value.String())
			}
			return carapace.ActionValues()
		}),
	})
}
