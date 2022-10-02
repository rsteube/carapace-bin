package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lightdm",
	Short: "a display manager",
	Long:  "https://github.com/canonical/lightdm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("cache-dir", "", "Directory to cache information")
	rootCmd.Flags().StringP("config", "c", "", "Use configuration file")
	rootCmd.Flags().BoolP("debug", "d", false, "Print debugging messages")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().String("log-dir", "", "Directory to write logs to")
	rootCmd.Flags().String("pid-file", "", "File to write PID into")
	rootCmd.Flags().String("run-dir", "", "Directory to store running state")
	rootCmd.Flags().Bool("show-config", false, "Show combined configuration")
	rootCmd.Flags().Bool("test-mode", false, "Run as unprivileged user, skipping things that require root access")
	rootCmd.Flags().BoolP("version", "v", false, "Show release version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir": carapace.ActionDirectories(),
		"config":    carapace.ActionFiles(),
		"log-dir":   carapace.ActionDirectories(),
		"pid-file":  carapace.ActionFiles(),
		"run-dir":   carapace.ActionDirectories(),
	})
}
