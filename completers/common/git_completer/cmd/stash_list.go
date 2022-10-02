package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stash_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the stash entries that you currently have",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_listCmd).Standalone()

	stashCmd.AddCommand(stash_listCmd)
}
