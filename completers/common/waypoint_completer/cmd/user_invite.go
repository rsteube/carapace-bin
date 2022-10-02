package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var user_inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite a user to join the Waypoint server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_inviteCmd).Standalone()

	user_inviteCmd.Flags().String("expires-in", "", "The duration until the token expires.")
	user_inviteCmd.Flags().String("username", "", "Invite a new user and provide a username hint.")

	addGlobalOptions(user_inviteCmd)

	userCmd.AddCommand(user_inviteCmd)
}
