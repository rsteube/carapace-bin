package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Recover data from the repository not referenced by snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recoverCmd).Standalone()
	rootCmd.AddCommand(recoverCmd)
}
