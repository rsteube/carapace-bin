package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_kubeconfig_adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Generate a kubeconfig file for the admin to use and for kubeadm itself",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeconfig_adminCmd).Standalone()
	init_phase_kubeconfig_adminCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_kubeconfig_adminCmd.Flags().Int32("apiserver-bind-port", 6443, "Port for the API Server to bind to.")
	init_phase_kubeconfig_adminCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_kubeconfig_adminCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_kubeconfig_adminCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_kubeconfig_adminCmd.Flags().String("kubeconfig-dir", "/etc/kubernetes", "The path where to save the kubeconfig file.")
	init_phase_kubeconfig_adminCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_kubeconfigCmd.AddCommand(init_phase_kubeconfig_adminCmd)

	carapace.Gen(init_phase_kubeconfig_adminCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":       carapace.ActionDirectories(),
		"config":         carapace.ActionFiles(),
		"kubeconfig-dir": carapace.ActionDirectories(),
	})
}
