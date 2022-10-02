package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cache_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List filenames of packages stored in the cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_listCmd).Standalone()

	cache_listCmd.Flags().String("format", "", "Select the output format")
	cacheCmd.AddCommand(cache_listCmd)

	carapace.Gen(cache_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("human", "abspath"),
	})
}
