package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Work with the system cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_cacheCmd).Standalone()

	pub_cacheCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.AddCommand(pub_cacheCmd)
}
