package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_checkExpirationCmd = &cobra.Command{
	Use:   "check-expiration",
	Short: "Check certificates expiration for a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_checkExpirationCmd).Standalone()
	certs_checkExpirationCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save the certificates")
	certs_checkExpirationCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_checkExpirationCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	certsCmd.AddCommand(certs_checkExpirationCmd)

	carapace.Gen(certs_checkExpirationCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
