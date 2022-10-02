package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listAllCmd = &cobra.Command{
	Use:   "list-all",
	Short: "List all packages available in repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listAllCmd).Standalone()

	rootCmd.AddCommand(listAllCmd)
}
