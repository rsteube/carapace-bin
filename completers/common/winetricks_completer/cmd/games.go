package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gamesCmd = &cobra.Command{
	Use:   "games",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamesCmd).Standalone()

	rootCmd.AddCommand(gamesCmd)
}
