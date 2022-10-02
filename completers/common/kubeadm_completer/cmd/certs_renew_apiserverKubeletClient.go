package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_renew_apiserverKubeletClientCmd = &cobra.Command{
	Use:   "apiserver-kubelet-client",
	Short: "Renew the certificate for the API server to connect to kubelet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renew_apiserverKubeletClientCmd).Standalone()
	certs_renew_apiserverKubeletClientCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save the certificates")
	certs_renew_apiserverKubeletClientCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_renew_apiserverKubeletClientCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	certs_renewCmd.AddCommand(certs_renew_apiserverKubeletClientCmd)

	carapace.Gen(certs_renew_apiserverKubeletClientCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
