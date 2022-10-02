package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List OpenFaaS functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	listCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	listCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	listCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	listCmd.Flags().BoolP("quiet", "q", false, "Quiet mode - print out only the function's ID")
	listCmd.Flags().String("sort", "name", "Sort the functions by \"name\" or \"invocations\"")
	listCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	listCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	listCmd.Flags().BoolP("verbose", "v", false, "Verbose output for the function list")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
		"sort":      carapace.ActionValues("name", "invocations"),
	})
}
