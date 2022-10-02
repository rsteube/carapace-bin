package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upgrade_node_phase_kubeletConfigCmd = &cobra.Command{
	Use:   "kubelet-config",
	Short: "Upgrade the kubelet configuration for this node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phase_kubeletConfigCmd).Standalone()
	upgrade_node_phase_kubeletConfigCmd.Flags().Bool("dry-run", false, "Do not change any state, just output the actions that would be performed.")
	upgrade_node_phase_kubeletConfigCmd.Flags().String("kubeconfig", "/etc/kubernetes/kubelet.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_node_phaseCmd.AddCommand(upgrade_node_phase_kubeletConfigCmd)

	carapace.Gen(upgrade_node_phase_kubeletConfigCmd).FlagCompletion(carapace.ActionMap{
		"kubeconfig": carapace.ActionFiles(),
	})
}
