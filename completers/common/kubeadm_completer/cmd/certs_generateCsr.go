package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_generateCsrCmd = &cobra.Command{
	Use:   "generate-csr",
	Short: "Generate keys and certificate signing requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_generateCsrCmd).Standalone()
	certs_generateCsrCmd.Flags().String("cert-dir", "", "The path where to save the certificates")
	certs_generateCsrCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_generateCsrCmd.Flags().String("kubeconfig-dir", "/etc/kubernetes", "The path where to save the kubeconfig file.")
	certsCmd.AddCommand(certs_generateCsrCmd)

	carapace.Gen(certs_generateCsrCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":       carapace.ActionDirectories(),
		"config":         carapace.ActionFiles(),
		"kubeconfig-dir": carapace.ActionDirectories(),
	})
}
