package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	secret_createCmd.Flags().String("from-file", "", "Path and filename containing value for the secret")
	secret_createCmd.Flags().String("from-literal", "", "Literal value for the secret")
	secret_createCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	secret_createCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	secret_createCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	secret_createCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	secret_createCmd.Flags().Bool("trim", true, "Trim whitespace from the start and end of the secret value")
	secretCmd.AddCommand(secret_createCmd)

	carapace.Gen(secret_createCmd).FlagCompletion(carapace.ActionMap{
		"from-file": carapace.ActionFiles(),
		"namespace": action.ActionNamespaces(),
	})

	carapace.Gen(secret_createCmd).PositionalCompletion(
		action.ActionSecrets(),
	)
}
