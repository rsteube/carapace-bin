package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "View authentication status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_statusCmd).Standalone()
	auth_statusCmd.Flags().StringP("hostname", "h", "", "Check a specific instance's authentication status")
	auth_statusCmd.Flags().BoolP("show-token", "t", false, "Display the auth token")
	authCmd.AddCommand(auth_statusCmd)

	carapace.Gen(auth_statusCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
	})
}
