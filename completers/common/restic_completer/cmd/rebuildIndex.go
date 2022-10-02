package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rebuildIndexCmd = &cobra.Command{
	Use:   "rebuild-index",
	Short: "Build a new index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebuildIndexCmd).Standalone()
	rebuildIndexCmd.Flags().Bool("read-all-packs", false, "read all pack files to generate new index from scratch")
	rootCmd.AddCommand(rebuildIndexCmd)
}
