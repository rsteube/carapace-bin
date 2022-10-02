package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listUnitFilesCmd = &cobra.Command{
	Use:   "list-unit-files",
	Short: "List installed unit files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listUnitFilesCmd).Standalone()

	rootCmd.AddCommand(listUnitFilesCmd)
}
