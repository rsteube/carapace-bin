package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the value of a given configuration key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()
	config_getCmd.Flags().StringP("host", "h", "", "Get per-host setting")
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionConfigHosts(),
	})

	carapace.Gen(config_getCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"browser", "the web browser to use for opening URLs",
			"git_protocol", "What protocol to use when performing git operations.",
			"editor", "What editor gh should run when creating issues, pull requests, etc.",
			"prompt", "toggle interactive prompting in the terminal",
			"pager", "the terminal pager program to send standard output to",
			"http_unix_socket", "the path to a unix socket through which to make HTTP connection",
		),
	)
}
