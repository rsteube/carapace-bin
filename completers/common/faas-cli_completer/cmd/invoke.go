package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "Invoke an OpenFaaS function",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	invokeCmd.Flags().BoolP("async", "a", false, "Invoke the function asynchronously")
	invokeCmd.Flags().String("content-type", "text/plain", "The content-type HTTP header such as application/json")
	invokeCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	invokeCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	invokeCmd.Flags().StringArrayP("header", "H", []string{}, "pass HTTP request header")
	invokeCmd.Flags().String("key", "", "key to be used to sign the request (must be used with --sign)")
	invokeCmd.Flags().StringP("method", "m", "POST", "pass HTTP request method")
	invokeCmd.Flags().String("name", "", "Name of the deployed function")
	invokeCmd.Flags().StringP("namespace", "n", "", "Namespace of the deployed function")
	invokeCmd.Flags().StringArray("query", []string{}, "pass query-string options")
	invokeCmd.Flags().String("sign", "", "name of HTTP request header to hold the signature")
	invokeCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	rootCmd.AddCommand(invokeCmd)

	carapace.Gen(invokeCmd).FlagCompletion(carapace.ActionMap{
		"content-type": http.ActionMediaTypes(),
		"method":       http.ActionRequestMethods(),
		"namespace":    action.ActionNamespaces(),
	})

	carapace.Gen(invokeCmd).PositionalCompletion(
		action.ActionFunctions(),
	)
}
