package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var join_phase_controlPlaneJoin_updateStatusCmd = &cobra.Command{
	Use:   "update-status",
	Short: "Register the new control-plane node into the ClusterStatus maintained in the kubeadm-config ConfigMap (DEPRECATED)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_controlPlaneJoin_updateStatusCmd).Standalone()
	join_phase_controlPlaneJoin_updateStatusCmd.Flags().String("apiserver-advertise-address", "", "If the node should host a new control plane instance, the IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	join_phase_controlPlaneJoin_updateStatusCmd.Flags().String("config", "", "Path to kubeadm config file.")
	join_phase_controlPlaneJoin_updateStatusCmd.Flags().Bool("control-plane", false, "Create a new control plane instance on this node")
	join_phase_controlPlaneJoin_updateStatusCmd.Flags().String("node-name", "", "Specify the node name.")
	join_phase_controlPlaneJoinCmd.AddCommand(join_phase_controlPlaneJoin_updateStatusCmd)

	carapace.Gen(join_phase_controlPlaneJoin_updateStatusCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
