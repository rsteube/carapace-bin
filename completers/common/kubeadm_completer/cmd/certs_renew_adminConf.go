package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_renew_adminConfCmd = &cobra.Command{
	Use:   "admin.conf",
	Short: "Renew the certificate embedded in the kubeconfig file for the admin to use and for kubeadm itself",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renew_adminConfCmd).Standalone()
	certs_renew_adminConfCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save the certificates")
	certs_renew_adminConfCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_renew_adminConfCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster.")
	certs_renewCmd.AddCommand(certs_renew_adminConfCmd)

	carapace.Gen(certs_renew_adminConfCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
