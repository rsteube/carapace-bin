package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listKeysCmd = &cobra.Command{
	Use:   "list-keys",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listKeysCmd).Standalone()

	listKeysCmd.Flags().BoolS("1", "1", false, "TODO description")
	listKeysCmd.Flags().BoolS("N", "N", false, "TODO description")
	listKeysCmd.Flags().StringS("P", "P", "", "prefix-string")
	listKeysCmd.Flags().StringS("T", "T", "", "key-table")
	listKeysCmd.Flags().BoolS("a", "a", false, "TODO description")
	rootCmd.AddCommand(listKeysCmd)
}
