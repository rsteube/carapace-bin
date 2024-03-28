package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var cluster_editCmd = &cobra.Command{
	Use:     "edit CLUSTER",
	Short:   "[EXPERIMENTAL] Edit cluster(s).",
	Aliases: []string{"update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_editCmd).Standalone()

	cluster_editCmd.Flags().StringSlice("port-add", []string{}, "[EXPERIMENTAL] Map ports from the node containers (via the serverlb) to the host (Format: `[HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER]`)")
	clusterCmd.AddCommand(cluster_editCmd)

	carapace.Gen(cluster_editCmd).FlagCompletion(carapace.ActionMap{
		"port-add": carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
			cluster := ""
			if len(c.Args) > 0 {
				cluster = c.Args[0]
			}

			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues()
			default:
				return k3d.ActionNodes(k3d.NodeOpts{Cluster: cluster, Running: true, Stopped: true}).UniqueList(",") // TODO nodefilter?
			}
		}),
	})

	carapace.Gen(cluster_editCmd).PositionalCompletion(
		k3d.ActionClusters(),
	)
}
