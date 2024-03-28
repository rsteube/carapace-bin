package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cluster_stopCmd = &cobra.Command{
	Use:   "stop [NAME [NAME...] | --all]",
	Short: "Stop existing k3d cluster(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_stopCmd).Standalone()

	cluster_stopCmd.Flags().BoolP("all", "a", false, "Stop all existing clusters")
	clusterCmd.AddCommand(cluster_stopCmd)
}
