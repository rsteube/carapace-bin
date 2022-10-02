package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_cache_repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "Reinstall cached packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_cache_repairCmd).Standalone()

	pub_cache_repairCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_cacheCmd.AddCommand(pub_cache_repairCmd)
}
