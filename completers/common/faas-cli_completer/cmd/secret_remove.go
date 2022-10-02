package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	secret_removeCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	secret_removeCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	secret_removeCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	secret_removeCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	secretCmd.AddCommand(secret_removeCmd)

	carapace.Gen(secret_removeCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
	})

	carapace.Gen(secret_removeCmd).PositionalCompletion(
		action.ActionSecrets(),
	)
}
