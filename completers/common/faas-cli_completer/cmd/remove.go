package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove deployed OpenFaaS functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	removeCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	removeCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	removeCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	removeCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	removeCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
	})
}
