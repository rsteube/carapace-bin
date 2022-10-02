package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var respawnWindowCmd = &cobra.Command{
	Use:   "respawn-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(respawnWindowCmd).Standalone()

	respawnWindowCmd.Flags().StringS("c", "c", "", "start-directory")
	respawnWindowCmd.Flags().StringS("e", "e", "", "environment")
	respawnWindowCmd.Flags().BoolS("k", "k", false, "TODO description")
	respawnWindowCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(respawnWindowCmd)
}
