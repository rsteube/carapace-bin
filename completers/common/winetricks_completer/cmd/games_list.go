package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var games_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list verbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(games_listCmd).Standalone()

	gamesCmd.AddCommand(games_listCmd)
}
