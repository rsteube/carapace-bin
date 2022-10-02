package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an Android Virtual Device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()

	rootCmd.AddCommand(deleteCmd)
}
