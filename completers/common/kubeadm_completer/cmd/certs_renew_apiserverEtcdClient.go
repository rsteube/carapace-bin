package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_renew_apiserverEtcdClientCmd = &cobra.Command{
	Use:   "apiserver-etcd-client",
	Short: "Renew the certificate the apiserver uses to access etcd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renew_apiserverEtcdClientCmd).Standalone()
	certs_renew_apiserverEtcdClientCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save the certificates")
	certs_renew_apiserverEtcdClientCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_renew_apiserverEtcdClientCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	certs_renewCmd.AddCommand(certs_renew_apiserverEtcdClientCmd)

	carapace.Gen(certs_renew_apiserverEtcdClientCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
