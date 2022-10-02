package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var server_configSetCmd = &cobra.Command{
	Use:   "config-set",
	Short: "Set the server online configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_configSetCmd).Standalone()

	server_configSetCmd.Flags().String("advertise-addr", "", "Address to advertise for the server.")
	server_configSetCmd.Flags().Bool("advertise-tls", false, "If true, the advertised address should be connected to with TLS.")
	server_configSetCmd.Flags().Bool("advertise-tls-skip-verify", false, "Do not verify the TLS certificate presented by the server. ")

	addGlobalOptions(server_configSetCmd)

	serverCmd.AddCommand(server_configSetCmd)
}
