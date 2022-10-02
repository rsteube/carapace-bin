package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listCachedCmd = &cobra.Command{
	Use:   "list-cached",
	Short: "list cached-and-ready-to-install verbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCachedCmd).Standalone()

	rootCmd.AddCommand(listCachedCmd)
}
