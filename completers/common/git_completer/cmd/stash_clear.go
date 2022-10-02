package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stash_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "remove all the stash entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_clearCmd).Standalone()

	stashCmd.AddCommand(stash_clearCmd)
}
