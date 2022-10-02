package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var teams_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show all teams you're a part of",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teams_lsCmd).Standalone()

	teamsCmd.AddCommand(teams_lsCmd)
}
