package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_renew_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Renew all available certificates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renew_allCmd).Standalone()
	certs_renew_allCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save the certificates")
	certs_renew_allCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_renew_allCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	certs_renewCmd.AddCommand(certs_renew_allCmd)

	carapace.Gen(certs_renew_allCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
