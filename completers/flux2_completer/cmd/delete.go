package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete sources and resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()
	deleteCmd.PersistentFlags().BoolP("silent", "s", false, "delete resource without asking for confirmation")
	rootCmd.AddCommand(deleteCmd)
}
