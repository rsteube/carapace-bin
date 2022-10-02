package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var join_phase_controlPlanePrepare_controlPlaneCmd = &cobra.Command{
	Use:   "control-plane",
	Short: "Generate the manifests for the new control plane components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_controlPlanePrepare_controlPlaneCmd).Standalone()
	join_phase_controlPlanePrepare_controlPlaneCmd.Flags().String("apiserver-advertise-address", "", "If the node should host a new control plane instance, the IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	join_phase_controlPlanePrepare_controlPlaneCmd.Flags().Int32("apiserver-bind-port", 6443, "If the node should host a new control plane instance, the port for the API Server to bind to.")
	join_phase_controlPlanePrepare_controlPlaneCmd.Flags().String("config", "", "Path to kubeadm config file.")
	join_phase_controlPlanePrepare_controlPlaneCmd.Flags().Bool("control-plane", false, "Create a new control plane instance on this node")
	join_phase_controlPlanePrepare_controlPlaneCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	join_phase_controlPlanePrepareCmd.AddCommand(join_phase_controlPlanePrepare_controlPlaneCmd)

	carapace.Gen(join_phase_controlPlanePrepare_controlPlaneCmd).FlagCompletion(carapace.ActionMap{
		"config":  carapace.ActionFiles(),
		"patches": carapace.ActionDirectories(),
	})
}
