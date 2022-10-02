package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upgrade_node_phase_controlPlaneCmd = &cobra.Command{
	Use:   "control-plane",
	Short: "Upgrade the control plane instance deployed on this node, if any",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phase_controlPlaneCmd).Standalone()
	upgrade_node_phase_controlPlaneCmd.Flags().Bool("certificate-renewal", true, "Perform the renewal of certificates used by component changed during upgrades.")
	upgrade_node_phase_controlPlaneCmd.Flags().Bool("dry-run", false, "Do not change any state, just output the actions that would be performed.")
	upgrade_node_phase_controlPlaneCmd.Flags().Bool("etcd-upgrade", true, "Perform the upgrade of etcd.")
	upgrade_node_phase_controlPlaneCmd.Flags().String("kubeconfig", "/etc/kubernetes/kubelet.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_node_phase_controlPlaneCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	upgrade_node_phaseCmd.AddCommand(upgrade_node_phase_controlPlaneCmd)

	carapace.Gen(upgrade_node_phase_controlPlaneCmd).FlagCompletion(carapace.ActionMap{
		"kubeconfig": carapace.ActionFiles(),
		"patches":    carapace.ActionDirectories(),
	})
}
