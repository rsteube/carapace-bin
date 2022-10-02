package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_renew_schedulerConfCmd = &cobra.Command{
	Use:   "scheduler.conf",
	Short: "Renew the certificate embedded in the kubeconfig file for the scheduler manager to use",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renew_schedulerConfCmd).Standalone()
	certs_renew_schedulerConfCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save the certificates")
	certs_renew_schedulerConfCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_renew_schedulerConfCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster.")
	certs_renewCmd.AddCommand(certs_renew_schedulerConfCmd)

	carapace.Gen(certs_renew_schedulerConfCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
