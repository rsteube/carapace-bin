package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var node_createCmd = &cobra.Command{
	Use:   "create NAME",
	Short: "Create a new k3s node in docker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_createCmd).Standalone()

	node_createCmd.Flags().StringP("cluster", "c", "", "Cluster URL or k3d cluster name to connect to.")
	node_createCmd.Flags().StringP("image", "i", "", "Specify k3s image used for the node(s) (default: copied from existing node)")
	node_createCmd.Flags().StringSlice("k3s-arg", []string{}, "Additional args passed to k3d command")
	node_createCmd.Flags().StringSlice("k3s-node-label", []string{}, "Specify k3s node labels in format \"foo=bar\"")
	node_createCmd.Flags().String("memory", "", "Memory limit imposed on the node [From docker]")
	node_createCmd.Flags().StringSliceP("network", "n", []string{}, "Add node to (another) runtime network")
	node_createCmd.Flags().String("replicas", "", "Number of replicas of this node specification.")
	node_createCmd.Flags().String("role", "", "Specify node role [server, agent]")
	node_createCmd.Flags().StringSlice("runtime-label", []string{}, "Specify container runtime labels in format \"foo=bar\"")
	node_createCmd.Flags().StringSlice("runtime-ulimit", []string{}, "Specify container runtime ulimit in format \"ulimit=soft:hard\"")
	node_createCmd.Flags().String("timeout", "", "Maximum waiting time for '--wait' before canceling/returning.")
	node_createCmd.Flags().StringP("token", "t", "", "Override cluster token (required when connecting to an external cluster)")
	node_createCmd.Flags().Bool("wait", false, "Wait for the node(s) to be ready before returning.")
	nodeCmd.AddCommand(node_createCmd)

	// TODO
	carapace.Gen(node_createCmd).FlagCompletion(carapace.ActionMap{
		// "cluster":        carapace.ActionValues(),
		// "image":          carapace.ActionValues(),
		"k3s-arg": bridge.ActionCarapaceBin("k3s").Split(), // TODO verify
		// "network":        carapace.ActionValues(),
		"role": carapace.ActionValues("server", "agent"),
	})
}
