package cmd

import (
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Fetch logs for a functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	logsCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	logsCmd.Flags().Bool("instance", false, "print the function instance name/id")
	logsCmd.Flags().Int("lines", -1, "number of recent log lines file to display. Defaults to -1, unlimited if <=0")
	logsCmd.Flags().Bool("name", false, "print the function name")
	logsCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	logsCmd.Flags().StringP("output", "o", "plain", "output logs as (plain|keyvalue|json), JSON includes all available keys")
	logsCmd.Flags().Duration("since", 0*time.Second, "return logs newer than a relative duration like 5s")
	logsCmd.Flags().String("since-time", "", "include logs since the given timestamp (RFC3339)")
	logsCmd.Flags().BoolP("tail", "t", true, "tail logs and continue printing new logs until the end of the request, up to 30s")
	logsCmd.Flags().String("time-format", "", "string format for the timestamp, any value go time format string is allowed, empty will not print the timestamp")
	logsCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	logsCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	rootCmd.AddCommand(logsCmd)

	carapace.Gen(logsCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
		"output":    carapace.ActionValues("plain", "keyvalue", "json"),
	})

	carapace.Gen(logsCmd).PositionalCompletion(
		action.ActionFunctions(),
	)
}
