package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var join_phase_controlPlaneJoin_etcdCmd = &cobra.Command{
	Use:   "etcd",
	Short: "Add a new local etcd member",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_controlPlaneJoin_etcdCmd).Standalone()
	join_phase_controlPlaneJoin_etcdCmd.Flags().String("apiserver-advertise-address", "", "If the node should host a new control plane instance, the IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	join_phase_controlPlaneJoin_etcdCmd.Flags().String("config", "", "Path to kubeadm config file.")
	join_phase_controlPlaneJoin_etcdCmd.Flags().Bool("control-plane", false, "Create a new control plane instance on this node")
	join_phase_controlPlaneJoin_etcdCmd.Flags().String("node-name", "", "Specify the node name.")
	join_phase_controlPlaneJoin_etcdCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	join_phase_controlPlaneJoinCmd.AddCommand(join_phase_controlPlaneJoin_etcdCmd)

	carapace.Gen(join_phase_controlPlaneJoin_etcdCmd).FlagCompletion(carapace.ActionMap{
		"config":  carapace.ActionFiles(),
		"patches": carapace.ActionDirectories(),
	})
}
