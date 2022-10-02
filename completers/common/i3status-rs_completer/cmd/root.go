package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "i3status-rs",
	Short: "A feature-rich and resource-friendly replacement for i3status, written in Rust",
	Long:  "https://github.com/greshake/i3status-rust",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("exit-on-error", false, "Exit rather than printing errors to i3bar and continuing")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().Bool("never-pause", false, "Ignore any attempts by i3 to pause the bar when hidden/fullscreen")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
