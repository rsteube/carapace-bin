package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clamd",
	Short: "an anti-virus daemon",
	Long:  "http://www.clamav.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config-file", "c", "", "Read configuration from FILE")
	rootCmd.Flags().Bool("debug", false, "Enable debug mode")
	rootCmd.Flags().BoolP("foreground", "F", false, "Run in foreground; do not daemonize")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().StringP("log", "l", "", "Log into FILE")
	rootCmd.Flags().BoolP("version", "V", false, "Show version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"log":         carapace.ActionFiles(),
	})
}
