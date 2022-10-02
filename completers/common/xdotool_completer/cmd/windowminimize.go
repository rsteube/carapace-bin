package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var windowminimizeCmd = &cobra.Command{
	Use:   "windowminimize",
	Short: "Minimize a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowminimizeCmd).Standalone()

	windowminimizeCmd.Flags().Bool("sync", false, "After requesting the window minimize, wait until the window is actually minimized")
	rootCmd.AddCommand(windowminimizeCmd)
}
