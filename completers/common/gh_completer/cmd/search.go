package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for repositories, issues, and pull requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()
	rootCmd.AddCommand(searchCmd)
}
