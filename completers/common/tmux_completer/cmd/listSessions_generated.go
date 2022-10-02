package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listSessionsCmd = &cobra.Command{
	Use:   "list-sessions",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listSessionsCmd).Standalone()

	listSessionsCmd.Flags().StringS("F", "F", "", "format")
	listSessionsCmd.Flags().StringS("f", "f", "", "filter")
	rootCmd.AddCommand(listSessionsCmd)
}
