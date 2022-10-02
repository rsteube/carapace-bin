package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Delete the specified cluster from the kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteClusterCmd).Standalone()

	configCmd.AddCommand(config_deleteClusterCmd)

	carapace.Gen(config_deleteClusterCmd).PositionalCompletion(
		action.ActionClusters(),
	)
}
