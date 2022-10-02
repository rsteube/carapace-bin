package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cache_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "delete all data out of the cache folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_cleanCmd).Standalone()

	cacheCmd.AddCommand(cache_cleanCmd)
}
