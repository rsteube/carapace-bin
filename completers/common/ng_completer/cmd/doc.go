package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Opens the official Angular documentation (angular.io) in a browser, and searches for a given keyword.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docCmd).Standalone()

	docCmd.Flags().BoolP("search", "s", false, "Search all of angular.io.")
	docCmd.Flags().String("version", "", "Contains the version of Angular to use for the documentation.")
	rootCmd.AddCommand(docCmd)
}
