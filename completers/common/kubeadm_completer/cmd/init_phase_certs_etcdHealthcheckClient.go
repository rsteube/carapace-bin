package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_etcdHealthcheckClientCmd = &cobra.Command{
	Use:   "etcd-healthcheck-client",
	Short: "Generate the certificate for liveness probes to healthcheck etcd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_etcdHealthcheckClientCmd).Standalone()
	init_phase_certs_etcdHealthcheckClientCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_certs_etcdHealthcheckClientCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_certs_etcdHealthcheckClientCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_certsCmd.AddCommand(init_phase_certs_etcdHealthcheckClientCmd)

	carapace.Gen(init_phase_certs_etcdHealthcheckClientCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
