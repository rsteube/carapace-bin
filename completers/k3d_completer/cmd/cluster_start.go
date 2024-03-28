package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cluster_startCmd = &cobra.Command{
	Use:   "start [NAME [NAME...] | --all]",
	Short: "Start existing k3d cluster(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_startCmd).Standalone()

	cluster_startCmd.Flags().BoolP("all", "a", false, "Start all existing clusters")
	cluster_startCmd.Flags().String("timeout", "", "Maximum waiting time for '--wait' before canceling/returning.")
	cluster_startCmd.Flags().Bool("wait", false, "Wait for the server(s) (and loadbalancer) to be ready before returning.")
	clusterCmd.AddCommand(cluster_startCmd)
}
