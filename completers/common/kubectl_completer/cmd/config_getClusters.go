package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_getClustersCmd = &cobra.Command{
	Use:   "get-clusters",
	Short: "Display clusters defined in the kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getClustersCmd).Standalone()

	configCmd.AddCommand(config_getClustersCmd)
}
