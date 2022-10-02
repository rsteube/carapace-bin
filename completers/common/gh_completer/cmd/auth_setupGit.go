package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_setupGitCmd = &cobra.Command{
	Use:   "setup-git",
	Short: "Configure git to use GitHub CLI as a credential helper",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_setupGitCmd).Standalone()
	auth_setupGitCmd.Flags().StringP("hostname", "h", "", "The hostname to configure git for")
	authCmd.AddCommand(auth_setupGitCmd)

	carapace.Gen(auth_setupGitCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
	})
}
