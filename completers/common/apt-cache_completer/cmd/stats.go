package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "stats displays some statistics about the cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()

	rootCmd.AddCommand(statsCmd)
}
