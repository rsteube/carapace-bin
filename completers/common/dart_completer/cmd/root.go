package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dart",
	Short: "A command-line utility for Dart development",
	Long:  "https://dart.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("disable-analytics", false, "Disable analytics.")
	rootCmd.Flags().Bool("enable-analytics", false, "Enable analytics.")
	rootCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.Flags().BoolP("verbose", "v", false, "Show additional command output.")
	rootCmd.Flags().Bool("version", false, "Print the Dart SDK version.")
}
