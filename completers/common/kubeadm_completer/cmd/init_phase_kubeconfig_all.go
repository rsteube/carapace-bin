package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_kubeconfig_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate all kubeconfig files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeconfig_allCmd).Standalone()
	init_phase_kubeconfig_allCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_kubeconfig_allCmd.Flags().Int32("apiserver-bind-port", 6443, "Port for the API Server to bind to.")
	init_phase_kubeconfig_allCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_kubeconfig_allCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_kubeconfig_allCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_kubeconfig_allCmd.Flags().String("kubeconfig-dir", "/etc/kubernetes", "The path where to save the kubeconfig file.")
	init_phase_kubeconfig_allCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_kubeconfig_allCmd.Flags().String("node-name", "", "Specify the node name.")
	init_phase_kubeconfigCmd.AddCommand(init_phase_kubeconfig_allCmd)

	carapace.Gen(init_phase_kubeconfig_allCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":       carapace.ActionDirectories(),
		"config":         carapace.ActionFiles(),
		"kubeconfig-dir": carapace.ActionDirectories(),
	})
}
