package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var team_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new user to an existing team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(team_addCmd).Standalone()

	teamCmd.AddCommand(team_addCmd)
}
