package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_secret_tlsCmd = &cobra.Command{
	Use:   "tls",
	Short: "Create or update a Kubernetes secret with TLS certificates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secret_tlsCmd).Standalone()
	create_secret_tlsCmd.Flags().String("ca-file", "", "TLS authentication CA file path")
	create_secret_tlsCmd.Flags().String("cert-file", "", "TLS authentication cert file path")
	create_secret_tlsCmd.Flags().String("key-file", "", "TLS authentication key file path")
	create_secretCmd.AddCommand(create_secret_tlsCmd)
}
