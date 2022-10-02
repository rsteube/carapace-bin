package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cache_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify the contents of the cache folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_verifyCmd).Standalone()

	cacheCmd.AddCommand(cache_verifyCmd)
}
