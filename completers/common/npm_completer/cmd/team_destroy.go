package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var team_destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroy an existing team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(team_destroyCmd).Standalone()

	teamCmd.AddCommand(team_destroyCmd)
}
