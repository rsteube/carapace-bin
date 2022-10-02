package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_addon_kubeProxyCmd = &cobra.Command{
	Use:   "kube-proxy",
	Short: "Install the kube-proxy addon to a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_addon_kubeProxyCmd).Standalone()
	init_phase_addon_kubeProxyCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_addon_kubeProxyCmd.Flags().Int32("apiserver-bind-port", 6443, "Port for the API Server to bind to.")
	init_phase_addon_kubeProxyCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_addon_kubeProxyCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_addon_kubeProxyCmd.Flags().String("image-repository", "k8s.gcr.io", "Choose a container registry to pull control plane images from")
	init_phase_addon_kubeProxyCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_addon_kubeProxyCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_addon_kubeProxyCmd.Flags().String("pod-network-cidr", "", "Specify range of IP addresses for the pod network. If set, the control plane will automatically allocate CIDRs for every node.")
	init_phase_addonCmd.AddCommand(init_phase_addon_kubeProxyCmd)

	carapace.Gen(init_phase_addon_kubeProxyCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
