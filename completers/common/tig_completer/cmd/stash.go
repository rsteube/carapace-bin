package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Start up in stash view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stashCmd).Standalone()

	rootCmd.AddCommand(stashCmd)
}
