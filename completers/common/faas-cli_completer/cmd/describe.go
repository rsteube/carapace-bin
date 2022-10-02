package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe an OpenFaaS function",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	describeCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	describeCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	describeCmd.Flags().String("name", "", "Name of the function")
	describeCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	describeCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	describeCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	rootCmd.AddCommand(describeCmd)

	carapace.Gen(describeCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
	})

	carapace.Gen(describeCmd).PositionalCompletion(
		action.ActionFunctions(),
	)
}
