package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/upower"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "upower",
	Short: "UPower command line tool",
	Long:  "https://upower.freedesktop.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("dump", "d", false, "Dump all parameters for all objects")
	rootCmd.Flags().BoolP("enumerate", "e", false, "Enumerate objects paths for devices")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().BoolP("monitor", "m", false, "Monitor activity from the power daemon")
	rootCmd.Flags().Bool("monitor-detail", false, "Monitor with detail")
	rootCmd.Flags().StringP("show-info", "i", "", "Show information about object path")
	rootCmd.Flags().BoolP("version", "v", false, "Print version of client and daemon")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"show-info": upower.ActionDevices(),
	})
}
