package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listWindowsCmd = &cobra.Command{
	Use:   "list-windows",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listWindowsCmd).Standalone()

	listWindowsCmd.Flags().StringS("F", "F", "", "format")
	listWindowsCmd.Flags().BoolS("a", "a", false, "TODO description")
	listWindowsCmd.Flags().StringS("f", "f", "", "filter")
	listWindowsCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(listWindowsCmd)
}
