package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_renew_controllerManagerConfCmd = &cobra.Command{
	Use:   "controller-manager.conf",
	Short: "Renew the certificate embedded in the kubeconfig file for the controller manager to use",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renew_controllerManagerConfCmd).Standalone()
	certs_renew_controllerManagerConfCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save the certificates")
	certs_renew_controllerManagerConfCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_renew_controllerManagerConfCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	certs_renewCmd.AddCommand(certs_renew_controllerManagerConfCmd)

	carapace.Gen(certs_renew_controllerManagerConfCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
