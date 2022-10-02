package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_apiserverEtcdClientCmd = &cobra.Command{
	Use:   "apiserver-etcd-client",
	Short: "Generate the certificate the apiserver uses to access etcd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_apiserverEtcdClientCmd).Standalone()
	init_phase_certs_apiserverEtcdClientCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_certs_apiserverEtcdClientCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_certs_apiserverEtcdClientCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_certsCmd.AddCommand(init_phase_certs_apiserverEtcdClientCmd)

	carapace.Gen(init_phase_certs_apiserverEtcdClientCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
