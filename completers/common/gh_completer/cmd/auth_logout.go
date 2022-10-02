package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of a GitHub host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_logoutCmd).Standalone()
	auth_logoutCmd.Flags().StringP("hostname", "h", "", "The hostname of the GitHub instance to log out of")
	authCmd.AddCommand(auth_logoutCmd)

	carapace.Gen(auth_logoutCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
	})
}
