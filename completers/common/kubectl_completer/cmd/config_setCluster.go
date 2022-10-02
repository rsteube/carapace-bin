package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setClusterCmd = &cobra.Command{
	Use:   "set-cluster",
	Short: "Sets a cluster entry in kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setClusterCmd).Standalone()

	config_setClusterCmd.Flags().String("embed-certs", "", "embed-certs for the cluster entry in kubeconfig")
	configCmd.AddCommand(config_setClusterCmd)
}
