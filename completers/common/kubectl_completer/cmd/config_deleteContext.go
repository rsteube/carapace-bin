package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_deleteContextCmd = &cobra.Command{
	Use:   "delete-context",
	Short: "Delete the specified context from the kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteContextCmd).Standalone()

	configCmd.AddCommand(config_deleteContextCmd)

	carapace.Gen(config_deleteContextCmd).PositionalCompletion(
		action.ActionContexts(),
	)
}
