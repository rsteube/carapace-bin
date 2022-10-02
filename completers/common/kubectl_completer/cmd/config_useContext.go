package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_useContextCmd = &cobra.Command{
	Use:   "use-context",
	Short: "Sets the current-context in a kubeconfig file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_useContextCmd).Standalone()

	configCmd.AddCommand(config_useContextCmd)

	carapace.Gen(config_useContextCmd).PositionalCompletion(
		action.ActionContexts(),
	)
}
