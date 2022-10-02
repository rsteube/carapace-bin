package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List centralized config entries of a given kind",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()

	config_listCmd.Flags().String("kind", "", "The kind of configurations to list.")
	config_listCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	configCmd.AddCommand(config_listCmd)

	carapace.Gen(config_listCmd).FlagCompletion(carapace.ActionMap{
		// TODO namespace
		"kind": action.ActionConfigKinds(),
	})
}
