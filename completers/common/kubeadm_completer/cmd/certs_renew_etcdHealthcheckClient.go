package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_renew_etcdHealthcheckClientCmd = &cobra.Command{
	Use:   "etcd-healthcheck-client",
	Short: "Renew the certificate for liveness probes to healthcheck etcd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renew_etcdHealthcheckClientCmd).Standalone()
	certs_renew_etcdHealthcheckClientCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save the certificates")
	certs_renew_etcdHealthcheckClientCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_renew_etcdHealthcheckClientCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	certs_renewCmd.AddCommand(certs_renew_etcdHealthcheckClientCmd)

	carapace.Gen(certs_renew_etcdHealthcheckClientCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
