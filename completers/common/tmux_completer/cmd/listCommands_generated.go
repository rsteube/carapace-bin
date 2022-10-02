package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listCommandsCmd = &cobra.Command{
	Use:   "list-commands",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCommandsCmd).Standalone()

	listCommandsCmd.Flags().StringS("F", "F", "", "format")
	rootCmd.AddCommand(listCommandsCmd)
}
