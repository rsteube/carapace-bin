package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cache_dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Show the cache directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_dirCmd).Standalone()

	cacheCmd.AddCommand(cache_dirCmd)
}
