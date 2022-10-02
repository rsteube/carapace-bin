package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List objects in the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalCompletion(
		carapace.ActionValues("blobs", "packs", "index", "snapshots", "keys", "locks"),
	)
}
