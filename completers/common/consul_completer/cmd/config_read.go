package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a centralized config entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_readCmd).Standalone()
	addClientFlags(config_readCmd)
	addServerFlags(config_readCmd)
	config_readCmd.Flags().String("kind", "", "The kind of configuration to delete.")
	config_readCmd.Flags().String("name", "", "The name of configuration to delete.")
	config_readCmd.Flags().String("namespace", "", "Specifies the namespace to query.")

	configCmd.AddCommand(config_readCmd)

	carapace.Gen(config_readCmd).FlagCompletion(carapace.ActionMap{
		"kind": action.ActionConfigKinds(),
		"name": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := config_readCmd.Flag("kind"); flag.Changed {
				return action.ActionConfigEntries(flag.Value.String())
			}
			return carapace.ActionValues()
		}),
	})
}
