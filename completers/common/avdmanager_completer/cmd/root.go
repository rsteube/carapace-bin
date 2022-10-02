package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "avdmanager",
	Short: "create and manage Android Virtual Devices",
	Long:  "https://developer.android.com/studio/command-line/avdmanager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("clear-cache", false, "Clear the SDK Manager repository manifest cache.")
	rootCmd.Flags().BoolP("help", "h", false, "Help on a specific command.")
	rootCmd.Flags().BoolP("silent", "s", false, "Silent mode, shows errors only.")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode, shows errors, warnings and all messages.")
}
