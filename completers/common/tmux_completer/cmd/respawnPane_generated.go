package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var respawnPaneCmd = &cobra.Command{
	Use:   "respawn-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(respawnPaneCmd).Standalone()

	respawnPaneCmd.Flags().StringS("c", "c", "", "start-directory")
	respawnPaneCmd.Flags().StringS("e", "e", "", "environment")
	respawnPaneCmd.Flags().BoolS("k", "k", false, "TODO description")
	respawnPaneCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(respawnPaneCmd)
}
