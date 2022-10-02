package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_frontProxyCaCmd = &cobra.Command{
	Use:   "front-proxy-ca",
	Short: "Generate the self-signed CA to provision identities for front proxy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_frontProxyCaCmd).Standalone()
	init_phase_certs_frontProxyCaCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_certs_frontProxyCaCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_certs_frontProxyCaCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_certsCmd.AddCommand(init_phase_certs_frontProxyCaCmd)

	carapace.Gen(init_phase_certs_frontProxyCaCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
