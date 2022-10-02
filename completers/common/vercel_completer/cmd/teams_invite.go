package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var teams_inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite a new member to a team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teams_inviteCmd).Standalone()

	teamsCmd.AddCommand(teams_inviteCmd)
}
