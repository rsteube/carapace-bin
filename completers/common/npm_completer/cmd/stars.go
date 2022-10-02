package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var starsCmd = &cobra.Command{
	Use:   "stars",
	Short: "View packages marked as favorites",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(starsCmd).Standalone()

	rootCmd.AddCommand(starsCmd)

	// TODO user completion
}
