package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "search for passwords files containing search-string when decryted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grepCmd).Standalone()

	rootCmd.AddCommand(grepCmd)
}
