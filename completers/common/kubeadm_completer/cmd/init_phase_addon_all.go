package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_addon_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Install all the addons",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_addon_allCmd).Standalone()
	init_phase_addon_allCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_addon_allCmd.Flags().Int32("apiserver-bind-port", 6443, "Port for the API Server to bind to.")
	init_phase_addon_allCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_addon_allCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_addon_allCmd.Flags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for various features.")
	init_phase_addon_allCmd.Flags().String("image-repository", "k8s.gcr.io", "Choose a container registry to pull control plane images from")
	init_phase_addon_allCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_addon_allCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_addon_allCmd.Flags().String("pod-network-cidr", "", "Specify range of IP addresses for the pod network. If set, the control plane will automatically allocate CIDRs for every node.")
	init_phase_addon_allCmd.Flags().String("service-cidr", "10.96.0.0/12", "Use alternative range of IP address for service VIPs.")
	init_phase_addon_allCmd.Flags().String("service-dns-domain", "cluster.local", "Use alternative domain for services, e.g. \"myorg.internal\".")
	init_phase_addonCmd.AddCommand(init_phase_addon_allCmd)

	carapace.Gen(init_phase_addon_allCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
