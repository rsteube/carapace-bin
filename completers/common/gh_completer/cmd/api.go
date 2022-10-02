package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Make an authenticated GitHub API request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiCmd).Standalone()
	apiCmd.Flags().Duration("cache", 0, "Cache the response, e.g. \"3600s\", \"60m\", \"1h\"")
	apiCmd.Flags().StringArrayP("field", "F", []string{}, "Add a typed parameter in `key=value` format")
	apiCmd.Flags().StringArrayP("header", "H", []string{}, "Add a HTTP request header in `key:value` format")
	apiCmd.Flags().String("hostname", "", "The GitHub hostname for the request (default \"github.com\")")
	apiCmd.Flags().BoolP("include", "i", false, "Include HTTP response status line and headers in the output")
	apiCmd.Flags().String("input", "", "The `file` to use as body for the HTTP request (use \"-\" to read from standard input)")
	apiCmd.Flags().StringP("jq", "q", "", "Query to select values from the response using jq syntax")
	apiCmd.Flags().StringP("method", "X", "GET", "The HTTP method for the request")
	apiCmd.Flags().Bool("paginate", false, "Make additional HTTP requests to fetch all pages of results")
	apiCmd.Flags().StringSliceP("preview", "p", []string{}, "GitHub API preview `names` to request (without the \"-preview\" suffix)")
	apiCmd.Flags().StringArrayP("raw-field", "f", []string{}, "Add a string parameter in `key=value` format")
	apiCmd.Flags().Bool("silent", false, "Do not print the response body")
	apiCmd.Flags().StringP("template", "t", "", "Format the response using a Go template")
	rootCmd.AddCommand(apiCmd)

	carapace.Gen(apiCmd).FlagCompletion(carapace.ActionMap{
		"header":   http.ActionRequestHeaders(),
		"hostname": action.ActionConfigHosts(),
		"input":    carapace.ActionFiles(),
		"method":   http.ActionRequestMethods(),
		"preview": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionApiPreviews().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(apiCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionApiV3Paths(apiCmd),
				carapace.ActionValues("graphql"),
			).Invoke(c).Merge().ToA()
		}),
	)
}
