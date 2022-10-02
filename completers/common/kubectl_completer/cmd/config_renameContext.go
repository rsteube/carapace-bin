package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_renameContextCmd = &cobra.Command{
	Use:   "rename-context",
	Short: "Renames a context from the kubeconfig file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_renameContextCmd).Standalone()

	configCmd.AddCommand(config_renameContextCmd)

	carapace.Gen(config_renameContextCmd).PositionalCompletion(
		action.ActionContexts(),
	)
}
