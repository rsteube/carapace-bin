package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cache_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove one or more package from the cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_removeCmd).Standalone()

	cacheCmd.AddCommand(cache_removeCmd)
}
