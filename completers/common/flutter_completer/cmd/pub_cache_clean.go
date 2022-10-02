package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_cache_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clears the global PUB_CACHE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_cache_cleanCmd).Standalone()

	pub_cache_cleanCmd.Flags().BoolP("force", "f", false, "Don't ask for confirmation.")
	pub_cache_cleanCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_cacheCmd.AddCommand(pub_cache_cleanCmd)
}
