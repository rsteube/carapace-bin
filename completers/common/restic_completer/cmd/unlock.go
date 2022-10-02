package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Remove locks other processes created",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlockCmd).Standalone()
	unlockCmd.Flags().Bool("remove-all", false, "remove all locks, even non-stale ones")
	rootCmd.AddCommand(unlockCmd)
}
