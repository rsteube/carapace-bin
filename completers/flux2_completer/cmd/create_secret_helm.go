package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_secret_helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Create or update a Kubernetes secret for Helm repository authentication",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secret_helmCmd).Standalone()
	create_secret_helmCmd.Flags().String("ca-file", "", "TLS authentication CA file path")
	create_secret_helmCmd.Flags().String("cert-file", "", "TLS authentication cert file path")
	create_secret_helmCmd.Flags().String("key-file", "", "TLS authentication key file path")
	create_secret_helmCmd.Flags().StringP("password", "p", "", "basic authentication password")
	create_secret_helmCmd.Flags().StringP("username", "u", "", "basic authentication username")
	create_secretCmd.AddCommand(create_secret_helmCmd)
}
