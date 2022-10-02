package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with a GitLab instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()
	auth_loginCmd.Flags().StringP("hostname", "h", "", "The hostname of the GitLab instance to authenticate with")
	auth_loginCmd.Flags().Bool("stdin", false, "Read token from standard input")
	auth_loginCmd.Flags().StringP("token", "t", "", "Your GitLab access token")
	authCmd.AddCommand(auth_loginCmd)

	carapace.Gen(auth_loginCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
	})
}
