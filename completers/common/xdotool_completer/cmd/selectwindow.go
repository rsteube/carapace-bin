package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var selectwindowCmd = &cobra.Command{
	Use:   "selectwindow",
	Short: "Get the window id (for a client) by clicking on it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectwindowCmd).Standalone()

	rootCmd.AddCommand(selectwindowCmd)
}
