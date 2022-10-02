package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Operate on local cache directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()
	cacheCmd.Flags().Bool("cleanup", false, "remove old cache directories")
	cacheCmd.Flags().Uint("max-age", 30, "max age in `days` for cache directories to be considered old")
	cacheCmd.Flags().Bool("no-size", false, "do not output the size of the cache directories")
	rootCmd.AddCommand(cacheCmd)
}
