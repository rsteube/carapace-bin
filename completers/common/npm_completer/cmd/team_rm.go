package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var team_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a user from a team they belong to",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(team_rmCmd).Standalone()

	teamCmd.AddCommand(team_rmCmd)
}
