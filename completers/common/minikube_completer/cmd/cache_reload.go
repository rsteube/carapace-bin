package cmd

import (
	"github.com/spf13/cobra"
)

var cache_reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "reload cached images.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cacheCmd.AddCommand(cache_reloadCmd)
}
