package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Stash a markdown",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stashCmd).Standalone()
	stashCmd.PersistentFlags().StringP("memo", "m", "", "memo/note for stashing")
	rootCmd.AddCommand(stashCmd)
}
