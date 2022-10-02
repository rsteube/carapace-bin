package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var renameSessionCmd = &cobra.Command{
	Use:   "rename-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameSessionCmd).Standalone()

	renameSessionCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(renameSessionCmd)
}
