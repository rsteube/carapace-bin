package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_setContextCmd = &cobra.Command{
	Use:   "set-context",
	Short: "Sets a context entry in kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setContextCmd).Standalone()

	config_setContextCmd.Flags().Bool("current", false, "Modify the current context")
	configCmd.AddCommand(config_setContextCmd)

	carapace.Gen(config_setContextCmd).PositionalCompletion(
		action.ActionContexts(),
	)
}
