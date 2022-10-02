package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_addon_corednsCmd = &cobra.Command{
	Use:   "coredns",
	Short: "Install the CoreDNS addon to a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_addon_corednsCmd).Standalone()
	init_phase_addon_corednsCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_addon_corednsCmd.Flags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for various features.")
	init_phase_addon_corednsCmd.Flags().String("image-repository", "k8s.gcr.io", "Choose a container registry to pull control plane images from")
	init_phase_addon_corednsCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_addon_corednsCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_addon_corednsCmd.Flags().String("service-cidr", "10.96.0.0/12", "Use alternative range of IP address for service VIPs.")
	init_phase_addon_corednsCmd.Flags().String("service-dns-domain", "cluster.local", "Use alternative domain for services, e.g. \"myorg.internal\".")
	init_phase_addonCmd.AddCommand(init_phase_addon_corednsCmd)
}
