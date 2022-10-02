package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "consul",
	Short: "Consul automates networking for simple and secure application delivery",
	Long:  "https://www.consul.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("version", "v", false, "show version")
}

func addClientFlags(cmd *cobra.Command) {
	cmd.Flags().String("ca-file", "", "Path to a CA file to use for TLS when communicating with Consul.")
	cmd.Flags().String("ca-path", "", "Path to a directory of CA certificates to use for TLS when communicating with Consul.")
	cmd.Flags().String("client-cert", "", "Path to a client cert file to use for TLS when 'verify_incoming' is enabled.")
	cmd.Flags().String("client-key", "", "Path to a client key file to use for TLS when 'verify_incoming' is enabled.")
	cmd.Flags().String("http-addr", "", "The `address` and port of the Consul HTTP agent.")
	cmd.Flags().String("tls-server-name", "", "The server name to use as the SNI host when connecting via TLS.")
	cmd.Flags().String("token", "", "ACL token to use in the request.")
	cmd.Flags().String("token-file", "", "File containing the ACL token to use in the request.")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"ca-file":     carapace.ActionFiles(),
		"ca-path":     carapace.ActionDirectories(),
		"client-cert": carapace.ActionFiles(),
		"client-key":  carapace.ActionFiles(),
		"token-file":  carapace.ActionFiles(),
	})
}

func addServerFlags(cmd *cobra.Command) {
	cmd.Flags().String("datacenter", "", "Name of the datacenter to query.")
	cmd.Flags().Bool("stale", false, "Permit any Consul server (non-leader) to respond to this request.")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"datacenter": action.ActionDatacenters(cmd),
	})
}

// TODO enable
//func addMultiTenancyFlags(cmd *cobra.Command) {
//	cmd.Flags().String("namespace", "", "Specifies the namespace to query.")
//
//	// TODO namespace completion
//}

// TODO enable
//func addPartitionFlags(cmd *cobra.Command) {
//	cmd.Flags().String("partition", "", "Specifies the admin partition to query.")
//
//	// TODO partition completion
//}
