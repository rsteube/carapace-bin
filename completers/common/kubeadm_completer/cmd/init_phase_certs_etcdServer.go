package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_etcdServerCmd = &cobra.Command{
	Use:   "etcd-server",
	Short: "Generate the certificate for serving etcd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_etcdServerCmd).Standalone()
	init_phase_certs_etcdServerCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_certs_etcdServerCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_certs_etcdServerCmd.Flags().String("kubernetes-version", "stable-1", "Choose a specific Kubernetes version for the control plane.")
	init_phase_certsCmd.AddCommand(init_phase_certs_etcdServerCmd)

	carapace.Gen(init_phase_certs_etcdServerCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
