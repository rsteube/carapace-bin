package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	secret_updateCmd.Flags().String("from-file", "", "Path to the secret file")
	secret_updateCmd.Flags().String("from-literal", "", "Value of the secret")
	secret_updateCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	secret_updateCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	secret_updateCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	secret_updateCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	secret_updateCmd.Flags().Bool("trim", true, "trim whitespace from the start and end of the secret value")
	secretCmd.AddCommand(secret_updateCmd)

	carapace.Gen(secret_updateCmd).FlagCompletion(carapace.ActionMap{
		"from-file": carapace.ActionFiles(),
		"namespace": action.ActionNamespaces(),
	})

	carapace.Gen(secret_updateCmd).PositionalCompletion(
		action.ActionSecrets(),
	)
}
