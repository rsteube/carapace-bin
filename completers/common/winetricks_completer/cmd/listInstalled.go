package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listInstalledCmd = &cobra.Command{
	Use:   "list-installed",
	Short: "list already-installed verbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listInstalledCmd).Standalone()

	rootCmd.AddCommand(listInstalledCmd)
}
