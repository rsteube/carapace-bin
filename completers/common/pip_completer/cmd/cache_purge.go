package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cache_purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Remove all items from the cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_purgeCmd).Standalone()

	cacheCmd.AddCommand(cache_purgeCmd)
}
