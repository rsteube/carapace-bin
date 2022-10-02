package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tls_cert_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new certificate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tls_cert_createCmd).Standalone()

	tls_cert_createCmd.Flags().String("additional-dnsname", "", "Provide an additional dnsname for Subject Alternative Names.")
	tls_cert_createCmd.Flags().String("additional-ipaddress", "", "Provide an additional ipaddress for Subject Alternative Names.")
	tls_cert_createCmd.Flags().String("ca", "", "Provide path to the ca.")
	tls_cert_createCmd.Flags().Bool("cli", false, "Generate cli certificate.")
	tls_cert_createCmd.Flags().Bool("client", false, "Generate client certificate.")
	tls_cert_createCmd.Flags().String("days", "", "Provide number of days the certificate is valid for from now on.")
	tls_cert_createCmd.Flags().String("dc", "", "Provide the datacenter.")
	tls_cert_createCmd.Flags().String("domain", "", "Provide the domain.")
	tls_cert_createCmd.Flags().String("key", "", "Provide path to the key.")
	tls_cert_createCmd.Flags().String("node", "", "When generating a server cert and this is set an additional dns name is included of the form <node>.server.<datacenter>.<domain>.")
	tls_cert_createCmd.Flags().Bool("server", false, "Generate server certificate.")
	tls_certCmd.AddCommand(tls_cert_createCmd)

	carapace.Gen(tls_cert_createCmd).FlagCompletion(carapace.ActionMap{
		// TODO flag completion
		"ca":  carapace.ActionFiles(),
		"key": carapace.ActionFiles(),
	})
}
