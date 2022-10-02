package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wineboot",
	Long:  "https://wiki.winehq.org/Wineboot",
	Short: "perform Wine initialization, startup, and shutdown task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("end-session", "e", false, "End the current session cleanly")
	rootCmd.Flags().BoolP("force", "f", false, "Force exit for processes that don't exit cleanly")
	rootCmd.Flags().BoolP("help", "h", false, "Display this help message")
	rootCmd.Flags().BoolP("init", "i", false, "Perform initialization for first Wine instance")
	rootCmd.Flags().BoolP("kill", "k", false, "Kill running processes without any cleanup")
	rootCmd.Flags().BoolP("restart", "r", false, "Restart only, don't do normal startup operations")
	rootCmd.Flags().BoolP("shutdown", "s", false, "Shutdown only, don't reboot")
	rootCmd.Flags().BoolP("update", "u", false, "Update the wineprefix directory")
}
