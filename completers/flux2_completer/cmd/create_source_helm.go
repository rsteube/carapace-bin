package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_source_helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Create or update a HelmRepository source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_source_helmCmd).Standalone()
	create_source_helmCmd.Flags().String("ca-file", "", "TLS authentication CA file path")
	create_source_helmCmd.Flags().String("cert-file", "", "TLS authentication cert file path")
	create_source_helmCmd.Flags().String("key-file", "", "TLS authentication key file path")
	create_source_helmCmd.Flags().Bool("pass-credentials", false, "pass credentials to all domains")
	create_source_helmCmd.Flags().StringP("password", "p", "", "basic authentication password")
	create_source_helmCmd.Flags().String("secret-ref", "", "the name of an existing secret containing TLS or basic auth credentials")
	create_source_helmCmd.Flags().String("url", "", "Helm repository address")
	create_source_helmCmd.Flags().StringP("username", "u", "", "basic authentication username")
	create_sourceCmd.AddCommand(create_source_helmCmd)
}
