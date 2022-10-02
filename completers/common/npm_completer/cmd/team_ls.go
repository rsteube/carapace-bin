package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var team_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list teams/users from organizations/teams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(team_lsCmd).Standalone()

	teamCmd.AddCommand(team_lsCmd)
}
