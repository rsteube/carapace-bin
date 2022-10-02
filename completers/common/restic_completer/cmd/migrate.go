package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply migrations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateCmd).Standalone()
	migrateCmd.Flags().BoolP("force", "f", false, "apply a migration a second time")
	rootCmd.AddCommand(migrateCmd)
}
