package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with a GitHub host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()
	auth_loginCmd.Flags().StringP("git-protocol", "p", "", "The protocol to use for git operations: {ssh|https}")
	auth_loginCmd.Flags().StringP("hostname", "h", "", "The hostname of the GitHub instance to authenticate with")
	auth_loginCmd.Flags().StringSliceP("scopes", "s", []string{}, "Additional authentication scopes to request")
	auth_loginCmd.Flags().BoolP("web", "w", false, "Open a browser to authenticate")
	auth_loginCmd.Flags().Bool("with-token", false, "Read token from standard input")
	authCmd.AddCommand(auth_loginCmd)

	carapace.Gen(auth_loginCmd).FlagCompletion(carapace.ActionMap{
		"git-protocol": carapace.ActionValues("ssh", "https"),
		"hostname":     action.ActionConfigHosts(),
		"scopes": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAuthScopes().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
