package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unbindKeyCmd = &cobra.Command{
	Use:   "unbind-key",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unbindKeyCmd).Standalone()

	unbindKeyCmd.Flags().StringS("T", "T", "", "key-table")
	unbindKeyCmd.Flags().BoolS("a", "a", false, "TODO description")
	unbindKeyCmd.Flags().BoolS("n", "n", false, "TODO description")
	unbindKeyCmd.Flags().BoolS("q", "q", false, "TODO description")
	rootCmd.AddCommand(unbindKeyCmd)
}
