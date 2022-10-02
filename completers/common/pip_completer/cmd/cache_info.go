package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cache_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show information about the cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_infoCmd).Standalone()

	cacheCmd.AddCommand(cache_infoCmd)
}
