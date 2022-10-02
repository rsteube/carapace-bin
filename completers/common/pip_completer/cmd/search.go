package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for PyPI packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringP("index", "i", "", "Base URL of Python Package Index (default")
	rootCmd.AddCommand(searchCmd)
}
