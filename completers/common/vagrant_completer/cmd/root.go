package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vagrant",
	Short: "tool for building and managing virtual machine environments",
	Long:  "https://www.vagrantup.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("color", false, "Enable color output")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug output")
	rootCmd.PersistentFlags().Bool("debug-timestamp", false, "Enable debug output with timestamps")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print this help")
	rootCmd.PersistentFlags().Bool("machine-readable", false, "Enable machine readable output")
	rootCmd.PersistentFlags().Bool("no-color", false, "Disable color output")
	rootCmd.PersistentFlags().Bool("no-tty", false, "Enable non-interactive output")
	rootCmd.PersistentFlags().Bool("timestamp", false, "Enable timestamps on log output")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Display Vagrant version")
}
