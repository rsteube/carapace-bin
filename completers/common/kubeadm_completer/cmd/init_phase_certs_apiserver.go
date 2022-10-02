package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_apiserverCmd = &cobra.Command{
	Use:   "apiserver",
	Short: "Generate the certificate for serving the Kubernetes API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_apiserverCmd).Standalone()
	init_phase_certs_apiserverCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_certs_apiserverCmd.Flags().StringSlice("apiserver-cert-extra-sans", []string{}, "Optional extra Subject Alternative Names (SANs) to use for the API Server serving certificate. Can be both IP addresses and DNS names.")
	init_phase_certs_apiserverCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_certs_apiserverCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_certs_apiserverCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_certs_apiserverCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_certs_apiserverCmd.Flags().String("service-cidr", "10.96.0.0/12", "Use alternative range of IP address for service VIPs.")
	init_phase_certs_apiserverCmd.Flags().String("service-dns-domain", "cluster.local", "Use alternative domain for services, e.g. \"myorg.internal\".")
	init_phase_certsCmd.AddCommand(init_phase_certs_apiserverCmd)

	carapace.Gen(init_phase_certs_apiserverCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
