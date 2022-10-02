package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var teams_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teams_addCmd).Standalone()

	teamsCmd.AddCommand(teams_addCmd)
}
