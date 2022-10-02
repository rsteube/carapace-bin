package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var team_createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(team_createCmd).Standalone()

	teamCmd.AddCommand(team_createCmd)
}
