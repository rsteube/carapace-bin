package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var cluster_deleteCmd = &cobra.Command{
	Use:     "delete [NAME [NAME ...] | --all]",
	Short:   "Delete cluster(s).",
	Aliases: []string{"del", "rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_deleteCmd).Standalone()

	cluster_deleteCmd.Flags().BoolP("all", "a", false, "Delete all existing clusters")
	cluster_deleteCmd.Flags().StringP("config", "c", "", "Path of a config file to use")
	clusterCmd.AddCommand(cluster_deleteCmd)

	carapace.Gen(cluster_deleteCmd).PositionalAnyCompletion(
		k3d.ActionClusters().FilterArgs(),
	)
}
