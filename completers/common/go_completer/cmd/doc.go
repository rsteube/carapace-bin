package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "show documentation for package or symbol",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docCmd).Standalone()

	docCmd.Flags().BoolS("all", "all", false, "Show all the documentation for the package")
	docCmd.Flags().BoolS("cmd", "cmd", false, "Treat a command like a regular package")
	docCmd.Flags().BoolS("short", "short", false, "One-line representation for each symbol")
	docCmd.Flags().BoolS("src", "src", false, "Show the full source code for the symbol")
	rootCmd.AddCommand(docCmd)
}
